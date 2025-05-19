package handler

import (
	"encoding/json"
	"net/http"

	"go-server/core/database"
	"go-server/core/model"
	"go-server/utils"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {

	const defaultUserRole = "owner"

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var userReq model.UserRegisterRequest
	err = json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.Loger.Info("User request body", "body", userReq)
	hashed, err := utils.HashPasssword(userReq.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx := db.MustBegin()
	dbRes, err := tx.Exec("INSERT INTO user VALUES(uuid(),?,?,?,?)", userReq.Name, defaultUserRole, userReq.Email, string(hashed))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, status := model.GenerateResponse(http.StatusCreated, "Successsfully created user", dbRes)

	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(jsonRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var loginReq model.UserLoginRequest
	err = json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	users := []model.User{}
	identifier := loginReq.Identifier

	err = db.Select(&users, "SELECT * FROM user WHERE name=? OR email=?", identifier, identifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	password := loginReq.Password
	err = utils.IsPasswordValid(users[0].Password, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, status := model.GenerateResponse(http.StatusOK, "Successfully Logged In", map[string]any{"token": "newToken"})

	w.WriteHeader(status)
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonRes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
