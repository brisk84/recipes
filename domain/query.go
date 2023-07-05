package domain

import (
	"recipes/api"
)

type Query struct {
	Ingredients  []string `json:"ingredients"`
	MaxTime      int      `json:"max_time"`
	SortByTime   string   `json:"sort_by_time"`
	MinRating    float64  `json:"min_rating"`
	SortByRating string   `json:"sort_by_rating"`
}

func (q *Query) FromApi(req api.Query) error {
	if req.Ingredients != nil {
		q.Ingredients = *req.Ingredients
	}
	if req.MaxTime != nil {
		q.MaxTime = *req.MaxTime
	}
	if req.SortByTime != nil {
		q.SortByTime = string(*req.SortByTime)
	}
	if req.MinRating != nil {
		q.MinRating = *req.MinRating
	}
	if req.SortByRating != nil {
		q.SortByRating = string(*req.SortByRating)
	}
	return nil
}
