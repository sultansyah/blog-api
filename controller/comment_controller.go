package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Comment interface {
	Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
}
