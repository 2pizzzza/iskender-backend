package user

import "context"

type UserRepository interface {
	GetUserByEmail(ctx context.Context, req *LoginUserDTO) (*User, error)
	CreateUser(ctx context.Context, req *RegisterUserDto) (*User, error)
}
