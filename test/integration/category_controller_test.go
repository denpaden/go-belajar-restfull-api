package integration

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/denpaden/go-belajar-restfull-api/app"
	"github.com/denpaden/go-belajar-restfull-api/app/excecption"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/domain"
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
	"github.com/denpaden/go-belajar-restfull-api/app/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_api_test")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func truncateDB(db *sql.DB) {
	db.Exec("TRUNCATE TABLE category")
}

func SetupRouter() http.Handler {
	db := SetupTestDB()
	validate := validator.New()
	router := httprouter.New()
	app.InitCategory(db, validate, router)
	app.InitProduct(db, validate, router)
	router.PanicHandler = excecption.ErrorHandler

	return middleware.NewAauthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	database := SetupTestDB()
	truncateDB(database)
	// setup router
	router := SetupRouter()
	//setup req body
	reqBody := strings.NewReader(`{"name":"Gadget"}`)
	// setup http api
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/go/api/category", reqBody)
	request.Header.Add("Content-Type", "application-json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	// setup recorder & call api
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// cek result status code dari response api
	response := recorder.Result()

	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 200, helper.GetResponseCode(responseBody))
	assert.Equal(t, "OK", helper.GetResponseStatus(responseBody))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	database := SetupTestDB()
	truncateDB(database)
	router := SetupRouter()

	reqBody := strings.NewReader(`{"name":""}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/go/api/category", reqBody)
	request.Header.Add("Content-Type", "application-json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 400, helper.GetResponseCode(responseBody))
	assert.Equal(t, "BAD REQUEST", helper.GetResponseStatus(responseBody))

}

func TestUpdateCategorySuccess(t *testing.T) {

	// create new data
	database := SetupTestDB()
	truncateDB(database)
	tx, _ := database.Begin()
	categoryRepository := category.NewCategoryRepository()
	save := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Minuman",
	})
	tx.Commit()

	//update data
	router := SetupRouter()
	reqBody := strings.NewReader(`{"name":"Senjata"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/go/api/category/"+strconv.Itoa(save.Id), reqBody)
	request.Header.Add("Content-Type", "application-json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 200, helper.GetResponseCode(responseBody))
	assert.Equal(t, "OK", helper.GetResponseStatus(responseBody))
	assert.Equal(t, save.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Senjata", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	// create new data
	database := SetupTestDB()
	truncateDB(database)
	tx, _ := database.Begin()
	categoryRepository := category.NewCategoryRepository()
	save := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Makanan",
	})
	tx.Commit()

	//update data
	router := SetupRouter()
	reqBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/go/api/category/"+strconv.Itoa(save.Id), reqBody)
	request.Header.Add("Content-Type", "application-json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 400, helper.GetResponseCode(responseBody))
	assert.Equal(t, "BAD REQUEST", helper.GetResponseStatus(responseBody))

}

func TestDeleteCategorySuccess(t *testing.T) {
	// create new data
	database := SetupTestDB()
	truncateDB(database)
	tx, _ := database.Begin()
	categoryRepository := category.NewCategoryRepository()
	save := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Minuman",
	})
	tx.Commit()

	//update data
	router := SetupRouter()
	reqBody := strings.NewReader(`{"name":"Senjata"}`)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/go/api/category/"+strconv.Itoa(save.Id), reqBody)
	request.Header.Add("Content-Type", "application-json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 200, helper.GetResponseCode(responseBody))
	assert.Equal(t, "OK", helper.GetResponseStatus(responseBody))

}

func TestDeleteCategoryUnAtuhorized(t *testing.T) {

	//update data
	router := SetupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/go/api/category/1", nil)
	request.Header.Add("Content-Type", "application-json")
	request.Header.Add("X-API-KEY", "s")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 401, helper.GetResponseCode(responseBody))
	assert.Equal(t, "UNAUTHORIZED", helper.GetResponseStatus(responseBody))

}

func TestDeleteCategoryNotFound(t *testing.T) {

}

func TestGetAllCategorySuccess(t *testing.T) {
	// create new data
	database := SetupTestDB()
	truncateDB(database)
	tx, _ := database.Begin()
	categoryRepository := category.NewCategoryRepository()
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Minuman",
	})
	tx.Commit()

	//update data
	router := SetupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/go/api/category", nil)
	request.Header.Add("X-API-KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 200, helper.GetResponseCode(responseBody))
	assert.Equal(t, "OK", helper.GetResponseStatus(responseBody))

}

func TestGetAllCategoryUnAuthorized(t *testing.T) {
	// create new data
	database := SetupTestDB()
	truncateDB(database)
	tx, _ := database.Begin()
	categoryRepository := category.NewCategoryRepository()
	categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Minuman",
	})
	tx.Commit()

	//update data
	router := SetupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/go/api/category", nil)
	request.Header.Add("X-API-KEY", "")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 401, helper.GetResponseCode(responseBody))
	assert.Equal(t, "UNAUTHORIZED", helper.GetResponseStatus(responseBody))
}

func TestGetByIdCategorySuccess(t *testing.T) {
	// create new data
	database := SetupTestDB()
	truncateDB(database)
	tx, _ := database.Begin()
	categoryRepository := category.NewCategoryRepository()
	save := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Minuman",
	})
	tx.Commit()

	//update data
	router := SetupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/go/api/category/"+strconv.Itoa(save.Id), nil)
	request.Header.Add("X-API-KEY", "RAHASIA")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 200, helper.GetResponseCode(responseBody))
	assert.Equal(t, "OK", helper.GetResponseStatus(responseBody))
	assert.Equal(t, save.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, save.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetByIdCategoryUnAuthorized(t *testing.T) {

	//update data
	router := SetupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/go/api/category/1", nil)
	request.Header.Add("X-API-KEY", " ")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	responseBody := helper.ReadResponseBody(response)
	fmt.Println(responseBody)

	assert.Equal(t, 401, helper.GetResponseCode(responseBody))
	assert.Equal(t, "UNAUTHORIZED", helper.GetResponseStatus(responseBody))

}
func TestGetByIdCategoryNotFound(t *testing.T) {

}
