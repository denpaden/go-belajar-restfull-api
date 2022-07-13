package middleware

import (
	"github.com/denpaden/go-belajar-restfull-api/app/feature/category/model/web"
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
	"net/http"
)

type AauthMiddleware struct {
	handler http.Handler
}

func NewAauthMiddleware(handler http.Handler) *AauthMiddleware {
	return &AauthMiddleware{handler: handler}
}

func (middleware *AauthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// simple cek auth, hanya mengecek headernya apakah ada X-AUTH-API dan isinya adalah RAHASIA, maka diperbolehkan untuk akses api
	if "RAHASIA" == request.Header.Get("X-API-KEY") {
		// ok
		// jika oke maka tinggal lanjutkan tulis response
		middleware.handler.ServeHTTP(writer, request)

	} else {
		// error
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		response := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.LoggerError(response.Status)
		helper.WriteToResponseBody(writer, response)
	}

}
