package main

import (
	"atom-project/internal/app"
	"atom-project/internal/db"
	"atom-project/internal/repository"
	"context"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	ctx := context.Background()

	store, err := db.New(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	service := app.New(repository.New(store))

	router := gin.Default()
	router.POST("/user_settings", service.SetSettings)
	router.POST("/user_lessons", service.GetUserLessons)
	router.Run("0.0.0.0:80")
}
