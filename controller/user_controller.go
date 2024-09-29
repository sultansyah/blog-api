package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	Profile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
