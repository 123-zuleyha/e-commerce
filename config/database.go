package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//LoadConfig dosyası , .env dosyasını yükler ve gerekli ayarları yapar.

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// ConnectDB veritabanı bağlantısı
func ConnectDB() {
	var err error

	dsn := os.Getenv("DB_DSN")
	DB, err = gorm.Open(postgres.Open(dsn) , &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v" , err)

	}

	log.Println("Database connection succesfully opened")
}
