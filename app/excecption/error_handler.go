package excecption

import (
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/web"
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// 404 : NOT FOUND
// 400 : BAD REQUEST
// 500 INTERNAL SERVER ERROR
//

func ErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {
	if notFoundError(writer, request, error) {
		return
	} else if validationErrors(writer, request, error) {
		return
	} else if internalServerError(writer, request, error) {
		return
	}
}

// 404 not found
func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error}

		helper.LoggerError(response.Status + ":" + exception.Error)
		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

// 400 bad request
func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		response := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error()}

		helper.LoggerError(response.Status + ":" + exception.Error())
		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}

// 500 internal server error
func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		response := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   exception.Error()}

		helper.LoggerError(response.Status + ":" + exception.Error())
		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}
