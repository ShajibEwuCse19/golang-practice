package services

import "ecom-project-improved/domain"

// service.go depend on port.go [Service and UserRepo] in the same directory
type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo, // now I need to implement the methods of Service interface here
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