package main

import (
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

	r := gin.Default()

	route.SetupAuthRouter(r, &cfg)
	route.SetupAuthMiddleWare(r, &cfg)
	route.SetupNotificationRouter(r, &cfg)
	route.SetupOrderRouter(r, &cfg)
	route.SetupStorageRouter(r, &cfg)
	route.SetupProductRouter(r, &cfg)
	r.Run(cfg.Port)
}
