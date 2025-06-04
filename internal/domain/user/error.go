package user

import "errors"

var (
	ErrUserNotFound     = errors.New("User not found")
	ErrUserAlreadyExist = errors.New("User with this email already exist")
)
