package handlers

import (
	"ecom-project-improved/database"
	"ecom-project-improved/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	var newUser database.User // instance/object of database User struct
	decoder := json.NewDecoder(r.Body) // decode user info
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" || newUser.Password == "" {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store() // created new user and store 
	utils.SendData(w, createdUser, http.StatusCreated)
	
}