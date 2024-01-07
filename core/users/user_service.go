// core/user/user_service.go

package users

import "errors"

type UserService interface {
	CreateUser(user User) error
	FindAllUsers() ([]User, error)
	FindUserById(id uint) (*User, error)
}

type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email cannot be empty")
	}

	if err := s.repo.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) FindAllUsers() ([]User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userServiceImpl) FindUserById(id uint) (*User, error) {
	if id == 0 {
		return nil, errors.New("invalid ID, must be greater than 0")
	}

	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}