package httpserver

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

type Server struct {
	App  *fiber.App
	Log  *zap.Logger
	host string
	port string
}

func New(log *zap.Logger, host, port string) *Server {
	app := fiber.New(fiber.Config{
		AppName:      "Iskender Backend",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,DELETE,PUT,PATCH,OPTIONS",
		AllowHeaders:     "*",
		ExposeHeaders:    "Content-Lengths",
		AllowCredentials: true,
	}))

	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		log.Info("Request handled",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", time.Since(start)),
		)
		return err
	})

	app.Use(recover.New())

	return &Server{
		App:  app,
		Log:  log,
		host: host,
		port: port,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	if err := s.App.Listen(addr); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.App.ShutdownWithContext(ctx); err != nil {
		s.Log.Error(fmt.Sprintf("error while closing HTTP server: %v", err))
	}
}
