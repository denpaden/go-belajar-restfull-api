package app

import (
	"database/sql"
	"github.com/denpaden/go-belajar-restfull-api/app/excecption"
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouterApp() *httprouter.Router {
	db := NewDB()
	validate := validator.New()
	router := httprouter.New()
	InitCategory(db, validate, router)
	InitProduct(db, validate, router)
	router.PanicHandler = excecption.ErrorHandler
	return router
}

func InitCategory(db *sql.DB, validate *validator.Validate, router *httprouter.Router) {
	categoryRepository := category.NewCategoryRepository()
	categoryService := category.NewCategoryService(categoryRepository, db, validate)
	categoryController := category.NewCategoryController(categoryService)

	router.GET("/go/api/category", categoryController.FindALl)
	router.GET("/go/api/category/:categoryId", categoryController.FindById)
	router.POST("/go/api/category", categoryController.Create)
	router.PUT("/go/api/category/:categoryId", categoryController.Update)
	router.DELETE("/go/api/category/:categoryId", categoryController.Delete)
}
func InitProduct(db *sql.DB, validate *validator.Validate, router *httprouter.Router) {

}
