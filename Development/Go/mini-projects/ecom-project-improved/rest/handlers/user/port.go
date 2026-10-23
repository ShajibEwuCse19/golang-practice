package user

import "ecom-project-improved/domain"

// Service -> business logic
type Service interface {
	Create(user domain.User) (*domain.User, error)
	List() ([]*domain.User, error)
	FindUserByEmailAndPassword(email string, password string) (*domain.User, error)
}