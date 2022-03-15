package main

import (
	"net/http"

	"github.com/flyrory/go-web-blog/app/http/middlewares"
	"github.com/flyrory/go-web-blog/bootstrap"
	"github.com/flyrory/go-web-blog/config"
	c "github.com/flyrory/go-web-blog/pkg/config"
	"github.com/flyrory/go-web-blog/pkg/logger"
	"github.com/flyrory/go-web-blog/pkg/route"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()
	route.SetRoute(router)

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
