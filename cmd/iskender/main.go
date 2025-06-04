package main

import (
	"fmt"

	"github.com/2pizzzza/IskenderBackend/internal/app"
	"github.com/2pizzzza/IskenderBackend/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Println("%w", err)
	}

	fmt.Println(cfg)
	app.New(cfg)
}
