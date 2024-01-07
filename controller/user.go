package controller

import "net/http"

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Find(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User find"))
}
