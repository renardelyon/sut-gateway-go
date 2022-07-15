package main

import (
	"context"
	"log"
	"sut-gateway-go/config"

	"sut-gateway-go/route"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("error when load config")
	}

	ctx := context.Background()

	r := gin.Default()

	route.SetupAuthRouter(ctx, r, &cfg)
	r.Run(cfg.Port)
}
