package main

import (
	"crudSvc/dbutils"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func postProductDetails(c *gin.Context) {
	type Payload struct {
		RowsPerPage int32 `json:"rowsPerPage"`
		Offset      int32 `json:"offset"`
	}
	var payload Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	prodDetails, totalProducts, err := dbutils.GetProductDetails(payload.RowsPerPage, payload.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products":      prodDetails,
		"totalProducts": totalProducts,
	})
}

func getProductDrawerDetails(c *gin.Context) {
	productName := c.Query("productName")
	productName, _ = url.QueryUnescape(productName)

	prodImage, errImg := GetImageByStringSerpApi(productName)
	prodDetails, errDetails := dbutils.GetProductDetailsByProductName(productName)

	if errImg != nil || errDetails != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch product"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"productImage":   prodImage,
			"productDetails": prodDetails,
		})
	}
}

func getHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(CORSMiddleware())
	router.POST("/products/details", postProductDetails)
	router.GET("/products/drawer", getProductDrawerDetails)
	router.GET("/hello", getHelloWorld)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
