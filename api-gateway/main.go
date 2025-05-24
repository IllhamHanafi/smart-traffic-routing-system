package main

import (
	"fmt"

	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/config"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/handler"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/internal"
	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/router"

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
