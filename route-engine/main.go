package main

import (
	"fmt"
	"route-engine/config"
	"route-engine/handler"
	"route-engine/internal"
	"route-engine/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// init handler and dependencies
	svc := internal.New()
	defer svc.Close()
	h := handler.New(svc)

	// Register routes
	router.RegisterRoute(r, h)

	// start server
	cfg := config.GetConfig()
	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
