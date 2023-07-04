package domain

import (
	"recipes/api"
)

type Query struct {
	Ingredients []string `json:"ingredients"`
	MaxTime     int      `json:"max_time"`
	SortByTime  string   `json:"sort"`
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
	return nil
}
