package user

import (

	"github.com/2pizzzza/IskenderBackend/internal/domain/user"
	"github.com/2pizzzza/IskenderBackend/internal/http/api"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{
	authService user.UserService
}

func NewUserHandler(authService user.UserService) *UserHandler{
	return &UserHandler{authService: authService}
}

func(uh *UserHandler) PortLogin(c *fiber.Ctx) error{
	var req api.LoginRequest

	if err := c.BodyParser(&req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid request"})
	}

	token, err := uh.authService.Login(c.Context(), req.Email, req.Password)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	return  c.JSON(api.LoginResponse{Token: &token})
}