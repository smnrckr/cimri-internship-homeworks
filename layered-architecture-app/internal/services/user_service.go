package services

import (
	"internal/models"
	"internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) GetUsers() []models.User {
	return s.userRepository.GetUsers()
}

func (s *UserService) GetUser(id int) (models.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) CreateUser(user models.UserCreateRequest) models.User {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) UpdateUser(id int, user models.UserUpdateRequest) (models.User, error) {
	return s.userRepository.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) {
	s.userRepository.DeleteUser(id)
}
