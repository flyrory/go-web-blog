package main

import (
	"net/http"

	"github.com/flyrory/go-web-blog/app/http/middlewares"
	"github.com/flyrory/go-web-blog/bootstrap"
	"github.com/flyrory/go-web-blog/pkg/logger"
	"github.com/flyrory/go-web-blog/pkg/route"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()
	route.SetRoute(router)

	err := http.ListenAndServe(":3009", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
