package test

import (
	"database/sql"
	"github.com/chwlr/golang-api/app"
	"github.com/chwlr/golang-api/controller"
	"github.com/chwlr/golang-api/helper"
	"github.com/chwlr/golang-api/middleware"
	"github.com/chwlr/golang-api/repository"
	"github.com/chwlr/golang-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_category_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler{
	db := setupTestDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}
func TestCreateCategoryFailed(t *testing.T) {


}
func TestUpdateCategorySuccess(t *testing.T) {

}
func TestUpdateCategoryFailed(t *testing.T) {

}
func TestGetCategorySuccess(t *testing.T) {

}
func TestGetCategoryFailed(t *testing.T) {

}
func TestDeleteCategorySuccess(t *testing.T) {

}
func TestDeleteCategoryFailed(t *testing.T) {

}
func TestListCategorySuccess(t *testing.T) {

}
func TestUnauthorized(t *testing.T) {

}