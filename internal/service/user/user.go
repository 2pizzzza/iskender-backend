package user

import (
	"context"
	"fmt"

	"github.com/2pizzzza/IskenderBackend/internal/domain/user"
	"github.com/2pizzzza/IskenderBackend/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{
	repo user.UserRepository
}

func NewUserRepository(repo user.UserRepository) *UserService{
	return &UserService{repo: repo}
}

func(us *UserService) Login(ctx context.Context, email, password string)(string, error){
	params := &user.LoginUserDTO{
		Email: email,
		Password: password,
	}
	res, err := us.repo.GetUserByUsernameAndPassword(ctx, params)
	if err != nil{
		return "", user.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil{
		return "", fmt.Errorf("incorrect password")
	}

	accessToken, err := utils.GenerateToken(int((res.Id)))
	if err != nil{
		return "", fmt.Errorf("failed generated access token: %v", err)
	}

	return accessToken, nil
}