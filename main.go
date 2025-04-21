package main

import (
	"github.com/Zhunisbekov/football_api/data"
	"github.com/Zhunisbekov/football_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Инициализация базы
	data.InitDB()

	// Роуты
	routes.RegisterMatchRoutes(r)

	// Запуск
	r.Run(":8080")
}
