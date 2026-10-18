package handlers

import (
	"ecom-project-improved/database"
	"ecom-project-improved/utils"
	"encoding/json"
	"net/http"
)

type RequestLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestLogin RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestLogin)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	user := database.FindUserByEmailAndPassword(requestLogin.Email, requestLogin.Password)
	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	utils.SendData(w, user, http.StatusCreated) 
}