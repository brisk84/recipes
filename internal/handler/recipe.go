package handler

import (
	"errors"
	"io"
	"net/http"
	"recipes/api"
	"recipes/domain"
	"time"
)

func (h *Handler) PostApiRecipeCCreate(w http.ResponseWriter, r *http.Request) {
	sd := r.Context().Value(domain.SessionDataKey)
	if sd == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, err := parseRequest[api.Recipe](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	var reqd domain.Recipe
	err = reqd.FromCreate(req)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	h.lg.Infoln("create", req)
	ret, err := h.uc.CreateRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeCUpdate(w http.ResponseWriter, r *http.Request) {
	sd := r.Context().Value(domain.SessionDataKey)
	if sd == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	req, err := parseRequest[api.RecipeWithId](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	var reqd domain.Recipe
	err = reqd.FromUpdate(req)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	h.lg.Infoln("update", req)
	ret, err := h.uc.UpdateRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeCDelete(w http.ResponseWriter, r *http.Request) {
	sd := r.Context().Value(domain.SessionDataKey)
	if sd == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	req, err := parseRequest[api.Id](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	var reqd domain.ID
	err = reqd.FromApi(req)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	h.lg.Infoln("delete", req)
	ret, err := h.uc.DeleteRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeQList(w http.ResponseWriter, r *http.Request) {
	h.lg.Infoln("list")
	ret, err := h.uc.ListRecipes(r.Context())
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeQRead(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest[api.Id](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	var reqd domain.ID
	err = reqd.FromApi(req)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	h.lg.Infoln("read", req)
	ret, err := h.uc.ReadRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeQFind(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest[api.Query](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	var reqd domain.Query
	err = reqd.FromApi(req)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	ret, err := h.uc.FindRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeCVote(w http.ResponseWriter, r *http.Request) {
	sd := r.Context().Value(domain.SessionDataKey)
	if sd == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	req, err := parseRequest[api.Vote](r)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	var reqd domain.Vote
	err = reqd.FromApi(req)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	if reqd.Mark < 1 || reqd.Mark > 5 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	h.lg.Infoln("vote", req)
	reqd.UserId = sd.(*domain.SessionData).Login
	reqd.CrDt = time.Now()
	err = h.uc.VoteRecipe(r.Context(), reqd)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateRecord) {
			w.WriteHeader(http.StatusConflict)
			return
		}
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, "OK", nil)
}

func (h *Handler) PostApiRecipeCUpload(w http.ResponseWriter, r *http.Request) {
	// sd := r.Context().Value(domain.SessionDataKey)
	// if sd == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }
	h.lg.Infoln("upload")
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	defer file.Close()
	reqd := domain.FileInfoUpload{
		Id:     r.FormValue("recipe_id"),
		Step:   r.FormValue("step"),
		Reader: file,
		Size:   handler.Size,
	}
	err = h.uc.UploadRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, "OK", nil)
}

func (h *Handler) PostApiRecipeQDownload(w http.ResponseWriter, r *http.Request) {
	h.lg.Infoln("download")
	reqd := domain.FileInfoDownload{
		Id:   r.FormValue("recipe_id"),
		Step: r.FormValue("step"),
	}
	reader, err := h.uc.DownloadRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = io.Copy(w, reader)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
}
