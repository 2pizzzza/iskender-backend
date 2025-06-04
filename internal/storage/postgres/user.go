package postgres

import (
	"context"
	"fmt"

	"github.com/2pizzzza/IskenderBackend/internal/domain/user"
)

type UserRepository struct {
	queries *Queries
}

func NewUserRepository(db DBTX) *UserRepository {
	return &UserRepository{
		queries: New(db),
	}
}

func (up *UserRepository) GetUserByEmail(ctx context.Context, req *user.LoginUserDTO) (*user.User, error) {
	rowUser, err := up.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	res := &user.User{
		Id:        int(rowUser.ID),
		Username:  rowUser.Username,
		Email:     rowUser.Email,
		CreatedAt: rowUser.CreatedAt.Time,
		Password:  rowUser.Password,
	}
	return res, nil
}

func (up *UserRepository) CreateUser(ctx context.Context, req *user.RegisterUserDto) (*user.User, error) {
	params := CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	rowUser, err := up.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed create user: %v", err)
	}

	res := &user.User{
		Id:       int(rowUser.ID),
		Username: rowUser.Username,
		Email:    rowUser.Email,
		Password: rowUser.Password,
	}

	return res, nil
}
