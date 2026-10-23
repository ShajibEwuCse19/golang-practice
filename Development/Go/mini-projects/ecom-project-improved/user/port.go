package user

import (
	"ecom-project-improved/domain"
	userHandler "ecom-project-improved/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embedding [copy the methods of userHandler.Service]
}

// Port -> contracts/interfaces (dependency inversion principle -> repostiory depends on ports not other way around)
// User domain depends on UserRepo interface (UML -> dependency arrow points toward dependency)
type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	List() ([]*domain.User, error)
	FindUserByEmailAndPassword(email string, password string) (*domain.User, error)
}