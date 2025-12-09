package database

import (
	"log"
	"os"
	"github.com/AliasgharHeidari/test-gorm/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	dsn := os.Getenv("dsn")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to db")

}

func AutoMigrate() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("error while migration")
	}
}

func GetDB() *gorm.DB {
	return DB
}
