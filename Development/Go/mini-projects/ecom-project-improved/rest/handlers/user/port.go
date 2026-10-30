package user

import "ecom-project-improved/domain"

// Service defines the business logic contract required by the user HTTP handler.
type Service interface {
	Create(user domain.User) (*domain.User, error)
	List() ([]*domain.User, error)
	FindUserByEmailAndPassword(email string, password string) (*domain.User, error)
}