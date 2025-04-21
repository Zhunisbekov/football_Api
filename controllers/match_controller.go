package controllers

import (
	"net/http"

	"github.com/Iliyas/football-api/data"
	"github.com/gin-gonic/gin"
)

func GetMatches(c *gin.Context) {
	c.JSON(http.StatusOK, data.Matches)
}
