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

func (up *UserRepository) GetUserByUsernameAndPassword(ctx context.Context, req *user.LoginUserDTO) (*user.User, error) {
	params := GetUserByUsernameAndPasswordParams{
		Email:    req.Email,
		Password: req.Password,
	}

	rowUser, err := up.queries.GetUserByUsernameAndPassword(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	res := &user.User{
		Id:       int(rowUser.ID),
		Username: rowUser.Username,
		Email:    rowUser.Email,
		CreatedAt: rowUser.CreatedAt.Time,
	}
	return res, nil
}
