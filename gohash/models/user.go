package models

import "testptr/interfaces"

type User struct {
	ID       int32
	Name     string
	Password string
	hasher   interfaces.PasswordHasher
}

func NewUser(id int32, name, password string, hasher interfaces.PasswordHasher) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: password,
		hasher:   hasher,
	}
}

func (u *User) HashPassword() string {
	return u.hasher.HashPassword(u.Password)
}
