package controller

import "net/http"

type UserController interface {
	CreateUser() http.HandlerFunc
	GetUser() http.HandlerFunc
}
