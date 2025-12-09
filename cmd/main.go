package main

import (
	"github.com/AliasgharHeidari/test-gorm/internal/api/server"
	"github.com/AliasgharHeidari/test-gorm/internal/database"
/* 		seeder "test-gorm/internal/seeder" */
)

func main() {
	database.ConnectDB()
	database.AutoMigrate()
/* 	seeder.SeedUser() */
	server.Start()
}