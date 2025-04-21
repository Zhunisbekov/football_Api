package main

import (
	"github.com/Iliyas/football-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterMatchRoutes(r)

	r.Run(":8080")
}
