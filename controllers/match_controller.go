package controllers

import (
	"net/http"
	"strconv"

	"github.com/Zhunisbekov/football_api/data"
	"github.com/Zhunisbekov/football_api/models"
	"github.com/gin-gonic/gin"
)

func CreateMatch(c *gin.Context) {
	var newMatch models.Match

	if err := c.ShouldBindJSON(&newMatch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Присваиваем новый ID
	newMatch.ID = len(data.Matches) + 1
	data.Matches = append(data.Matches, newMatch)

	c.JSON(http.StatusCreated, newMatch)
}
func GetMatches(c *gin.Context) {
	c.JSON(http.StatusOK, data.Matches)
}

func GetMatchByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	for _, match := range data.Matches {
		if match.ID == id {
			c.JSON(http.StatusOK, match)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Матч не найден"})
}

func UpdateMatch(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var updatedMatch models.Match
	if err := c.ShouldBindJSON(&updatedMatch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, match := range data.Matches {
		if match.ID == id {
			updatedMatch.ID = id
			data.Matches[i] = updatedMatch
			c.JSON(http.StatusOK, updatedMatch)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Матч не найден"})
}

func DeleteMatch(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	for i, match := range data.Matches {
		if match.ID == id {
			data.Matches = append(data.Matches[:i], data.Matches[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Матч удалён"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Матч не найден"})
}
