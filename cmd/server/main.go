package main

import (
	"log"

	"github.com/ScMofeoluwa/GatherGo/config"
	"github.com/ScMofeoluwa/GatherGo/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Config error: %s", err)
	}

	app.Run(cfg)
}
