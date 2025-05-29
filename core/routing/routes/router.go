package routes

import (
	"github.com/gorilla/mux"

	"go-server/core/routing/handler"
	"go-server/core/routing/middleware"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(middleware.ApplicationJsonMiddleware)
	global := r.PathPrefix("/api").Subrouter()
	auth := global.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", handler.HandleLogin).Methods("POST")
	auth.HandleFunc("/register", handler.HandleRegister).Methods("POST")

	guard := global.NewRoute().Subrouter()
	guard.Use(middleware.AuthGuardJWTMiddleware)
	guard.HandleFunc("/create/post", handler.HandleCreatePost).Methods("POST")
}
