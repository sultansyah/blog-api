package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sultansyah/blog-api/service"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{service: service}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}

func (controller *UserControllerImpl) Profile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
