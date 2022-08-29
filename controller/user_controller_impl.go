package controller

import (
	"github.com/chwlr/golang-api/helper"
	"github.com/chwlr/golang-api/model/web"
	"github.com/chwlr/golang-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(
		request.Context(),
		userCreateRequest,
	)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := param.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userId := param.ByName("userId")
	id, err := strconv.Atoi(userId)
	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}
	helper.PanicIfError(err)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userId := param.ByName("userId")
	id, err := strconv.Atoi(userId)
	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}
	helper.PanicIfError(err)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	userResponses := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
