package middleware

import (
	"github.com/chwlr/golang-api/helper"
	"github.com/chwlr/golang-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "SECRET" == request.Header.Get("X-API-Key"){
		middleware.Handler.ServeHTTP(writer, request)
	}else {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
