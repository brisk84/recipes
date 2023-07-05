package domain

import (
	"recipes/api"
	"time"
)

type Vote struct {
	CrDt     time.Time `json:"cr_dt"`
	RecipeId string    `json:"recipe_id"`
	UserId   string    `json:"user_id"`
	Mark     int       `json:"mark"`
}

func (v *Vote) FromApi(req api.Vote) error {
	v.RecipeId = req.RecipeId
	v.Mark = req.Mark
	return nil
}
