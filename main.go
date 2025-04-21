package main

import (
	"github.com/Zhunisbekov/football-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterMatchRoutes(r)

	r.Run(":8080")
}
