package routes

import (
	"github.com/gorilla/mux"

	"go-server/core/routing/handler"

)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/login", handler.HandleLogin).Methods("POST")
	r.HandleFunc("/auth/register", handler.HandleRegister).Methods("POST")
}
