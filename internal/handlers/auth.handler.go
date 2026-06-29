package handlers

import (
	"errors"
	"net/http"
	"time"
	"github.com/DarkDeity666/Go_APIs/internal/database"
	"github.com/DarkDeity666/Go_APIs/internal/models"
	"github.com/DarkDeity666/Go_APIs/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Register(c *gin.Context) {

	// Create a variable to store the incoming request body
	var req models.RegisterRequest

	// Read JSON body and fill the req variable
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	// Get the users collection from MongoDB
	collection := database.Client.
		Database("GoAPIs").
		Collection("users")

	// Empty struct where MongoDB will store MongoDB result
	var existingUser models.User

	// Search user by email
	err := collection.FindOne(
		c.Request.Context(),
		bson.M{
			"email": req.Email,
		},
	).Decode(&existingUser)

	// User already exists
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "User already exists",
		})
		return
	}

	// Database error -- this check if there was error in database connection
	// or any other dataase error 
	if !errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Database error",
		})
		return
	}

	// Hash user's password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to hash password",
		})
		return
	}

	// Create User struct -- basically this is inserting data into User Struct
	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		Avatar:    "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// will insert struct into database form go local memeory 
	result ,err := collection.InsertOne(
		c.Request.Context(),
		user,
	)

	// this will check if the user was able to get craeted of not if now it throw err
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success":false,
			"message":"Failed to Create User",
		})
		return
	}
	_= result
	// Will return this if the user Registration is successfull
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User Registered successfully...!",
	})
}