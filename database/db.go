package database

import (
	"log"
	"os"

	"github.com/JuanCampbsi/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDataBase() {
	stringConnect := os.Getenv("STRING_CONNECTDB")
	DB, err = gorm.Open(postgres.Open(stringConnect))

	if err != nil {
		log.Panic("Erro connect database")
	}

	DB.AutoMigrate(&models.Student{})
}
