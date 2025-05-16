package routes

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/login", HandleLogin).Methods("POST")
	r.HandleFunc("/auth/register", HandleRegister).Methods("POST")
}
