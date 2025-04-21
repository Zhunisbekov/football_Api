package routes

import (
	"github.com/Zhunisbekov/football-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMatchRoutes(r *gin.Engine) {
	r.GET("/matches", controllers.GetMatches)
}
