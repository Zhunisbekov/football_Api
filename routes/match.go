package routes

import (
	"github.com/Iliyas/football-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMatchRoutes(r *gin.Engine) {
	r.GET("/matches", controllers.GetMatches)
}
