package main

import (
	"fmt"
	"log"

	"github.com/arsalanaa44/vote-api/initializers"
	"github.com/arsalanaa44/vote-api/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Poll{}, &models.Vote{})
	fmt.Println("? Migration complete")
}
