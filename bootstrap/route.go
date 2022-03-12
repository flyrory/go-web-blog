package bootstrap

import (
	"github.com/flyrory/go-web-blog/pkg/routers"
	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routers.RegisterWebRoutes(router)
	return router
}
