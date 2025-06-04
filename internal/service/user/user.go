package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/2pizzzza/IskenderBackend/internal/domain/user"
	"github.com/2pizzzza/IskenderBackend/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) Login(ctx context.Context, email, password string) (string, error) {
	params := &user.LoginUserDTO{
		Email:    email,
		Password: password,
	}
	res, err := us.repo.GetUserByEmail(ctx, params)
	if err != nil {
		fmt.Println(err)
		return "", user.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)); err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("incorrect password")
	}

	accessToken, err := utils.GenerateToken(int((res.Id)))
	if err != nil {
		return "", fmt.Errorf("failed generated access token: %v", err)
	}

	return accessToken, nil
}

func (us *UserService) Register(ctx context.Context, email, username, password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed hashed password: %v", err)
	}
	params := &user.RegisterUserDto{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	res, err := us.repo.CreateUser(ctx, params)
	if err != nil {
		if errors.Is(err, user.ErrUserAlreadyExist) {
			return "", user.ErrUserAlreadyExist
		}
		return "", fmt.Errorf("failed create user: %v", err)
	}

	accessToken, err := utils.GenerateToken(int((res.Id)))
	if err != nil {
		return "", fmt.Errorf("failed generated access token: %v", err)
	}

	return accessToken, nil
}
