package routes

import (
	"encoding/json"
	"net/http"

	"go-server/core/model"
	"go-server/utils"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	utils.Loger.Info("User request body", "body", user)
	res := model.GenerateResponse(model.Success, "Successsfully created user", user)
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

}
