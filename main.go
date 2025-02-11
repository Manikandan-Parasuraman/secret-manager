package main

import (
	"log"

	"github.com/Manikandan-Parasuraman/secret-manager/src/config"
	"github.com/Manikandan-Parasuraman/secret-manager/src/handlers"
	"github.com/Manikandan-Parasuraman/secret-manager/src/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	storage.ConnectDB()

	r := gin.Default()

	r.POST("/secret", handlers.CreateSecret)
	r.GET("/secret/:id", handlers.GetSecret)

	log.Println("Server started on port 8080")
	r.Run(":8080")
}
