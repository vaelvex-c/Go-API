package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
}

