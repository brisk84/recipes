package domain

import (
	"recipes/api"
	"time"
)

type ID struct {
	Id string `json:"id"`
}

func (i *ID) FromApi(req api.Id) error {
	i.Id = req.Id.String()
	return nil
}

type Base struct {
	CrDt   time.Time  `json:"-"`
	UpdDt  time.Time  `json:"-"`
	DelDt  *time.Time `json:"-"`
	UserId string     `json:"-"`
}
