package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var AppName = "NIR"
var AppVersion = "dev"    // Default value, overridden during build
var BuildTime = "unknown" // Default value, overridden during build

var userDBPath = "/Users/sofianurobert/Projects/invetory-app/api_server/CRUD/db/users.json"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func appInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"appName":    AppName,
		"appVersion": AppVersion,
		"buildTime":  BuildTime,
	})
}

func login(c *gin.Context) {
	// Parse the username and password from the request body
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println("Received login request:", input)

	// Open the user database file
	file, err := os.Open(userDBPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open user database"})
		return
	}
	defer file.Close()

	// Decode the user database
	var users []User
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user database"})
		return
	}

	// Check if the username and password match any user in the database
	for _, user := range users {
		if user.Username == input.Username && user.Password == input.Password {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			return
		}
	}

	// If no match is found, return an error
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}

func homeDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the home details page!",
	})
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default()) // Allows all origins, methods, and headers

	router.GET("/appInfo", appInfo)

	router.POST("/login", login)

	router.GET("/homeDetails", homeDetails)

	router.Run(":8080")
}
