package user

import "context"

type UserRepository interface {
	GetUserByUsernameAndPassword(ctx context.Context, req *LoginUserDTO) (*User, error)
}
