package app

import (
	"fmt"

	"github.com/2pizzzza/IskenderBackend/internal/config"
	"github.com/2pizzzza/IskenderBackend/pkg/logger"
)

func New(config *config.Config) {
	log, err := logger.New(config)
	if err != nil {
		fmt.Println("%w", err)
	}
	log.Info("Logger successful")
}
