package middlewares

import (
	"net/http"

	"github.com/flyrory/go-web-blog/pkg/auth"
	"github.com/flyrory/go-web-blog/pkg/flash"
)

// Auth 登录用户才可访问
func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !auth.Check() {
			flash.Warning("登录用户才能访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next(w, r)
	}
}
