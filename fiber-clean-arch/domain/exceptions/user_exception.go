package exceptions

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrDuplicatedEmail = errors.New("email already exist")
	ErrLoginFailed     = errors.New("login failed")
)
