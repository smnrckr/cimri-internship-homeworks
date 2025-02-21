package repositories

import (
	"errors"
	"internal/models"
)

type UserRepository struct {
	usersMap map[int]models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		usersMap: map[int]models.User{
			1: {Id: 1, Name: "John", Age: 25},
			2: {Id: 2, Name: "Jane", Age: 30},
		},
	}
}

func (r *UserRepository) GetUsers() []models.User {
	users := []models.User{}
	for _, user := range r.usersMap {
		users = append(users, user)
	}
	return users
}

func (r *UserRepository) GetUser(id int) (models.User, error) {
	user, ok := r.usersMap[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) CreateUser(userRequest models.UserCreateRequest) models.User {
	newUser := models.User{}
	newUser.Id = len(r.usersMap) + 1

	newUser.Name = userRequest.Name
	newUser.Age = userRequest.Age

	r.usersMap[newUser.Id] = newUser
	return newUser
}

func (r *UserRepository) UpdateUser(id int, updatedUser models.UserUpdateRequest) (models.User, error) {
	user, ok := r.usersMap[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}

	user.Name = updatedUser.Name
	user.Age = updatedUser.Age

	r.usersMap[id] = user

	return user, nil
}

func (r *UserRepository) DeleteUser(id int) {
	delete(r.usersMap, id)
}
