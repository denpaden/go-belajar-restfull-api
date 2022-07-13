package main

import (
	"github.com/denpaden/go-belajar-restfull-api/app"
	"github.com/denpaden/go-belajar-restfull-api/app/helper"
	"github.com/denpaden/go-belajar-restfull-api/app/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	router := app.NewRouterApp()
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAauthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
