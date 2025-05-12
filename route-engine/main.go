package main

import (
	"fmt"
	"route-engine/config"
	"route-engine/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes
	router.RegisterRoute(r)
	cfg := config.ParseConfig()

	r.Run(fmt.Sprintf(":%d", cfg.Port))
}
