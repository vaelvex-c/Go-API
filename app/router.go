package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/vaelvex/Go-API/controller"
	"github.com/vaelvex/Go-API/exception"
)

func NewRouter(categoryController controller.CategoryController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
