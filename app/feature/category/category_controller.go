package category

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindALl(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}