package user

import "context"

type UserService interface{
	Login(ctx context.Context, email, password string)(string, error)
}
