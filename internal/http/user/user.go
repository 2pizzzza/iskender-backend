package user

import (
	"fmt"

	"github.com/2pizzzza/IskenderBackend/api/user"
	"github.com/2pizzzza/IskenderBackend/internal/domain/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	authService user.UserService
}

func NewUserHandler(authService user.UserService) *UserHandler {
	return &UserHandler{authService: authService}
}

func (uh *UserHandler) PostLogin(c *fiber.Ctx) error {
	var req api.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	token, err := uh.authService.Login(c.Context(), string(req.Email), req.Password)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.JSON(api.LoginResponse{
		Token: &token,
	})
}

func (uh *UserHandler) PostRegister(c *fiber.Ctx) error {
	var req api.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Email == "" || req.Password == "" || req.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and username and password are required",
		})
	}

	token, err := uh.authService.Register(c.Context(), string(req.Email), req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.JSON(api.RegisterResponse{
		Token: &token,
	})
}
