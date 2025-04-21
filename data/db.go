package data

import (
	"log"

	"github.com/Zhunisbekov/football_api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("football.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Авто-миграция модели Match
	err = DB.AutoMigrate(&models.Match{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
}
