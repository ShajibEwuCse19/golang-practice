package user

import (
	"ecom-project-improved/config"
	"ecom-project-improved/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requestLogin RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestLogin)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.FindUserByEmailAndPassword(requestLogin.Email, requestLogin.Password)
	if err != nil || user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// create JWT token
	accessToken, err := utils.CreateJwt(
		config.GetConfig().JWTSecretKey,
		utils.PayLoad{
			Sub:         strconv.Itoa(int(user.ID)),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
		},
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusCreated)
}
