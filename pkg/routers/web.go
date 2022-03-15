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
	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	//创建文章页面
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	//修改页面
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	//修改文章
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	//删除文章
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))

	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 用户认证
	auc := new(controllers.AuthController)
	//用户注册
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")
	//用户登录
	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	//用户退出
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	// 用户相关
	uc := new(controllers.UserController)
	r.HandleFunc("/users/{id:[0-9]+}", uc.Show).Methods("GET").Name("users.show")
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	r.Use(middlewares.StartSession)
}
