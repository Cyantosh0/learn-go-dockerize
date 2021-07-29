package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/cyantosh0/dockerize-go-app/config"
	"github.com/cyantosh0/dockerize-go-app/route"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.DB = config.SetupDatabase()
	// config.DB.AutoMigrate(&model.User{})

	r := route.SetupRouter()

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
