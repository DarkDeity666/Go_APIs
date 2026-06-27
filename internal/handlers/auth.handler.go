package handlers

import (
	"errors"
	"net/http"

	"github.com/DarkDeity666/Go_APIs/internal/database"
	"github.com/DarkDeity666/Go_APIs/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Register(c *gin.Context) {

// func Register(c *gin.Context){
// 	// this create a variable named req which can hold the incoming request Body
// 	//also model.RegisterRequest it telling this variable that this is your datatype
// 	// and you have to store data in this formate
	var req models.RegisterRequest

// 	//this checks for the rquest Body and fill the req varible with the incoming request Body
// 	// and it return either err or nil
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	// Get users collection
	collection := database.Client.
		Database("GoAPIs").
		Collection("users")

	// Empty user to store MongoDB result
	var existingUser models.User

	// Search user by email
	err := collection.FindOne(
		c.Request.Context(),
		bson.M{
			"email": req.Email,
		},
	).Decode(&existingUser)

	// User not found -> Continue registration
	if errors.Is(err, mongo.ErrNoDocuments) {

		// We'll continue here in the next lesson

	} else if err != nil {

		// Some database error occurred
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Database error",
		})
		return

	} else {

		// User already exists
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "User already exists",
		})
		return
	}

	// Temporary response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Email is available. Ready to create user.",
	})
}