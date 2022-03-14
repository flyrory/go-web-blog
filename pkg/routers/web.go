package routers

import (
	"net/http"

	"github.com/flyrory/go-web-blog/app/http/controllers"
	"github.com/flyrory/go-web-blog/app/http/middlewares"
	"github.com/gorilla/mux"
)

func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pc := new(controllers.PagesController)
	//关于我们
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	ac := new(controllers.ArticlesController)
	//主页
	r.HandleFunc("/", ac.Index).Methods("GET").Name("home")
	//文章详情
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	//文章列表
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	//保存文章
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")
	//创建文章页面
	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	//修改页面
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	//修改文章
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	//删除文章
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))

	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 用户认证
	auc := new(controllers.AuthController)
	//用户注册
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")
	//用户登录
	r.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", auc.DoLogin).Methods("POST").Name("auth.dologin")
	//用户退出
	r.HandleFunc("/auth/logout", auc.Logout).Methods("POST").Name("auth.logout")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	r.Use(middlewares.StartSession)
}
