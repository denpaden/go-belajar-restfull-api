package category

import (
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/web"
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService CategoryService
}

//func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//
//	// decode data dari request
//	decoder := json.NewDecoder(r.Body)
//	categoryWebCreateRequest := web.CategoryWebCreateRequest{}
//	err := decoder.Decode(&categoryWebCreateRequest)
//	helper.PanicIfError(err)
//
//	// call serivice function
//	categoryWebResponse := controller.CategoryService.Save(r.Context(), categoryWebCreateRequest)
//	webResponse := web.ToWebResponse(200, "OK", categoryWebResponse)
//
//	// set header
//	w.Header().Add("Content-Type", "application-json")
//	encoder := json.NewEncoder(w)
//	err = encoder.Encode(webResponse)
//	helper.PanicIfError(err)
//}

func NewCategoryController(categoryService CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}
func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// read request
	categoryWebCreateRequest := web.CategoryWebCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryWebCreateRequest)

	// call serivice
	categoryWebResponse := controller.CategoryService.Save(r.Context(), categoryWebCreateRequest)
	webResponse := web.ToWebResponse(200, "OK", categoryWebResponse)

	// write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// read request
	categoryWebUpdateRequest := web.CategoryWebUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryWebUpdateRequest)

	// get id from param
	stringId := params.ByName("categoryId")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfError(err)
	categoryWebUpdateRequest.Id = id

	// call serivice function
	categoryWebResponse := controller.CategoryService.Update(r.Context(), categoryWebUpdateRequest)
	webResponse := web.ToWebResponse(200, "OK", categoryWebResponse)

	// write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// get id from param
	stringId := params.ByName("categoryId")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfError(err)

	// call serivice function
	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.ToWebResponseNoData(200, "OK")

	// write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// get id from param
	stringId := params.ByName("categoryId")
	id, err := strconv.Atoi(stringId)
	helper.PanicIfError(err)

	// call serivice function
	response := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.ToWebResponse(200, "OK", response)

	// write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindALl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// call serivice function
	response := controller.CategoryService.FindAll(r.Context())
	webResponse := web.ToWebResponse(200, "OK", response)

	// write response
	helper.WriteToResponseBody(w, webResponse)
}
