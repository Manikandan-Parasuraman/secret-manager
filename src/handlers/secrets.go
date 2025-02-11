package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Manikandan-Parasuraman/secret-manager/src/models"
	"github.com/Manikandan-Parasuraman/secret-manager/src/services"
	"github.com/Manikandan-Parasuraman/secret-manager/src/storage"
)

// Create a secret
func CreateSecret(c *gin.Context) {
	var secret models.Secret
	if err := c.BindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	encrypted, err := services.EncryptSecret(secret.Secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
		return
	}

	secret.Secret = encrypted
	secret.ID = primitive.NewObjectID().Hex()

	collection := storage.GetCollection("secrets")
	_, err = collection.InsertOne(context.Background(), secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store secret"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Secret stored successfully", "id": secret.ID})
}

// Retrieve a secret
func GetSecret(c *gin.Context) {
	id := c.Param("id")
	collection := storage.GetCollection("secrets")

	var secret models.Secret
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&secret)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Secret not found"})
		return
	}

	decrypted, err := services.DecryptSecret(secret.Secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": secret.Name, "secret": decrypted})
}
