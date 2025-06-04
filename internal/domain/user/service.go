package user

import "context"

type UserService interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, username, password string) (string, error)
}
