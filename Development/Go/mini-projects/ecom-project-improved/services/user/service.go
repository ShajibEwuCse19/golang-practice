package services

import "ecom-project-improved/domain"

// service implements the business logic for the user domain and depends on UserRepo interface.
type service struct {
	userRepo UserRepo
}

// NewService initializes a new user service with its repository dependency.
func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(user domain.User) (*domain.User, error) {
	usr, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

func (s *service) FindUserByEmailAndPassword(email string, password string) (*domain.User, error) {
	usr, err := s.userRepo.FindUserByEmailAndPassword(email,password)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

func (s *service) List() ([]*domain.User, error) {
	usr, err := s.userRepo.List()
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}