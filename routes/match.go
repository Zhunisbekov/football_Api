package routes

import (
	"github.com/Zhunisbekov/football_api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMatchRoutes(r *gin.Engine) {
	r.GET("/matches", controllers.GetMatches)
	r.GET("/matches/:id", controllers.GetMatchByID)
	r.POST("/matches", controllers.CreateMatch)
	r.PUT("/matches/:id", controllers.UpdateMatch)
	r.DELETE("/matches/:id", controllers.DeleteMatch)
}
