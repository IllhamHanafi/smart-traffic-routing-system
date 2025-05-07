package main

import (
	"github.com/gin-gonic/gin"
	"route-engine/router"
)

func main() {
	r := gin.Default()

	// Register routes
	router.RegisterPingRoute(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
