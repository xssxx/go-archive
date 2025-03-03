package services

import (
	"testptr/models"
)

type UserService struct {
	users map[int32]*models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int32]*models.User),
	}
}

func (us *UserService) AddUser(user *models.User) {
	us.users[user.ID] = user
}

func (us *UserService) ValidatePassword(userID int32, password string) bool {
	user, exists := us.users[userID]
	if !exists {
		return false
	}

	hashedPassword := user.HashPassword()

	return hashedPassword == password
}
