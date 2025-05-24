package main

import (
	"fmt"

	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/config"
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/handler"
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/internal"
	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/router"

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
