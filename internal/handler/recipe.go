package handler

import (
	"net/http"
	"recipes/api"
	"recipes/domain"
)

func (h *Handler) PostApiRecipeCCreate(w http.ResponseWriter, r *http.Request) {
	sd := r.Context().Value("SessionData")
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
	ret, err := h.uc.CreateRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeCUpdate(w http.ResponseWriter, r *http.Request) {
	sd := r.Context().Value("SessionData")
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
	ret, err := h.uc.UpdateRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeCDelete(w http.ResponseWriter, r *http.Request) {
	sd := r.Context().Value("SessionData")
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
	ret, err := h.uc.DeleteRecipe(r.Context(), reqd)
	if err != nil {
		h.lg.Errorln(err)
		sendResponse[NilType](w, nil, err)
		return
	}
	sendResponse(w, ret, nil)
}

func (h *Handler) PostApiRecipeQList(w http.ResponseWriter, r *http.Request) {
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
