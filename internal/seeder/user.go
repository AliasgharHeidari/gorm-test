package seeder

import (
	"log"
	"github.com/AliasgharHeidari/test-gorm/internal/database"
	"github.com/AliasgharHeidari/test-gorm/internal/model"
)

func SeedUser() {
	DB := database.GetDB()

	var user model.User
	user.UserName = "Ali2929"
	user.Email = "ali@gmail.com"
	user.Password = "12344321"
	
	err := DB.Create(&user).Error
	if err != nil {
		log.Fatal("failed to add user to database")
	}
}
