package domain

import (
	"recipes/api"
	"time"
)

type Recipe struct {
	Base
	ID
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}

func (r *Recipe) FromCreate(req api.Recipe) error {
	r.CrDt = time.Now()
	r.Title = req.Title
	r.Description = req.Description
	r.Ingredients = req.Ingredients
	r.Steps = req.Steps
	return nil
}

func (r *Recipe) FromUpdate(req api.RecipeWithId) error {
	r.CrDt = time.Now()
	r.Id = req.Id.String()
	r.Title = req.Title
	r.Description = req.Description
	r.Ingredients = req.Ingredients
	r.Steps = req.Steps
	return nil
}
