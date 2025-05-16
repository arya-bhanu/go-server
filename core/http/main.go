package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go-server/core/database"
	"go-server/core/routes"
	"go-server/utils"
)

func main() {
	const port = ":9000"

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	db, err := database.ConnectDB()
	if err != nil {
		utils.Loger.Error(err.Error())
		return
	}

	if db != nil {
		utils.Loger.Info("Database Connected Successfully")
	}

	server := http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	utils.Loger.Info("Listen and serve server at: ", "port", port)
	if err := server.ListenAndServe(); err != nil {
		utils.Loger.Error(err.Error())
	}

}
