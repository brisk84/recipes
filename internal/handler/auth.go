package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"recipes/api"
	"recipes/domain"
	"recipes/pkg/tools"
	"regexp"
	"time"
)

func (h *Handler) PostApiUserCSignin(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest[api.User](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	var reqd domain.User
	err = reqd.FromApi(req)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	ret, ok, err := h.uc.SignIn(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	if !ok {
		h.lg.Infoln("Invalid password", req.Login)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	sd := domain.SessionData{
		Login: ret.Login,
		Token: "Bearer " + tools.GetGuid(),
	}
	buf, err := json.Marshal(sd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	err = h.rc.Set(r.Context(), h.getKey(sd.Token), buf, time.Hour*12).Err()
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	w.Header().Add("Authorization", "Bearer "+h.getKey(sd.Token))
	res := map[string]interface{}{"auth_token": sd.Token}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sd := domain.SessionData{}
		ah := r.Header.Get("Authorization")
		sd.Token = h.getKey(ah)
		if sd.Token != "" {
			b, err := h.rc.Get(r.Context(), sd.Token).Bytes()
			if err == nil {
				err = json.Unmarshal(b, &sd)
				if err == nil {
					ctx := context.WithValue(r.Context(), domain.SessionDataKey, &sd)
					r = r.WithContext(ctx)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) getKey(s string) string {
	if s == "" {
		return ""
	}
	tpl := regexp.MustCompile("Bearer ([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$")
	if tpl.Match([]byte(s)) {
		matchs := tpl.FindAllString(s, -1)
		return matchs[0][7:]
	}
	return ""
}
