package routers

import (
	"github.com/gorilla/mux"
)

// Init 初始化路由信息
func Init() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
}
