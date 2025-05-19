package routes

import (
	"github.com/gorilla/mux"

	"go-server/core/routing/handler"
	"go-server/core/routing/middleware"

)

func RegisterRoutes(r *mux.Router) {
	r.Use(middleware.ApplicationJsonMiddleware)
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", handler.HandleLogin).Methods("POST")
	auth.HandleFunc("/register", handler.HandleRegister).Methods("POST")
}
