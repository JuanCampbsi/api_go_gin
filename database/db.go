package database

import (
	"log"
	"os"

	"github.com/JuanCampbsi/api-go-gin/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDataBase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stringConnect := os.Getenv("STRING_CONNECTDB")
	DB, err = gorm.Open(postgres.Open(stringConnect))

	if err != nil {
		log.Panic("Erro connect database")
	}

	DB.AutoMigrate(&models.Student{})
}
