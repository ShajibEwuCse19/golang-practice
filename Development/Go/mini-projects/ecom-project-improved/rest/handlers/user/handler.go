package user

import "ecom-project-improved/repo"

type Handler struct {
	userRepo repo.UserRepo
}

func NewHandler(userRepo repo.UserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}

