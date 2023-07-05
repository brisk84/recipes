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
	Steps       []Step   `json:"steps"`
	TotalTime   int      `json:"total_time"`
}

type RecipeForList struct {
	ID
	Title string `json:"title"`
}

type Step struct {
	Title string `json:"title"`
	Time  int    `json:"time"`
}

func (r *Recipe) FromCreate(req api.Recipe) error {
	r.CrDt = time.Now()
	r.Title = req.Title
	r.Description = req.Description
	r.Ingredients = req.Ingredients
	for _, v := range req.Steps {
		r.Steps = append(r.Steps, Step{
			Title: v.Title,
			Time:  v.Time,
		})
	}
	r.TotalTime = req.TotalTime
	return nil
}

func (r *Recipe) FromUpdate(req api.RecipeWithId) error {
	r.CrDt = time.Now()
	r.Id = req.Id.String()
	r.Title = req.Title
	r.Description = req.Description
	r.Ingredients = req.Ingredients
	for _, v := range req.Steps {
		r.Steps = append(r.Steps, Step{
			Title: v.Title,
			Time:  v.Time,
		})
	}
	r.TotalTime = req.TotalTime
	return nil
}
