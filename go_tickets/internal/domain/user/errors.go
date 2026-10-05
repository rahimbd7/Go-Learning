package user

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("user with this email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
)