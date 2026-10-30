package services

import (
	"ecom-project-improved/domain"
	userHandler "ecom-project-improved/rest/handlers/user"
)

// Service embeds the user handler contract required by the REST layer.
type Service interface {
	userHandler.Service
}

// UserRepo defines the storage contract (port) that the repository layer must implement.
// Dependency Inversion Principle: The service defines the port; the repository implements it.
type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	List() ([]*domain.User, error)
	FindUserByEmailAndPassword(email string, password string) (*domain.User, error)
}