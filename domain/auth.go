package domain

import "recipes/api"

type SessionData struct {
	Token string `json:"token"`
	Login string `json:"login"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *User) FromApi(req api.User) error {
	u.Login = req.Login
	u.Password = req.Password
	return nil
}
