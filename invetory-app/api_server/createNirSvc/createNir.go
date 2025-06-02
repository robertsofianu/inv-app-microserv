package main

import (
	"crudSvc/dbutils"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Username = "test"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ProductNameAmount struct {
	Product string `json:"product"`
	Amount  int32  `json:"amount"`
}

type NirForm struct {
	Company       string `json:"company" form:"company"`
	Adress        string `json:"adress" form:"adress"`
	NumberNIR     string `json:"nir_number" form:"nir_number"`
	DateNIR       string `json:"nir_date" form:"nir_date"`
	NumberInvoice string `json:"invoice_number" form:"invoice_number"`
	DateInvoice   string `json:"invoice_date" form:"invoice_date"`
	Provider      string `json:"provider" form:"provider"`
	Delegat       string `json:"delegat" form:"delegat"`
}

func nirCreate(c *gin.Context) {
	var form NirForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	CreateNirFile(form, form.NumberNIR)

	// Now form contains the submitted data
	c.JSON(http.StatusOK, gin.H{
		"company":        form.Company,
		"adress":         form.Adress,
		"delegat":        form.Delegat,
		"nir_number":     form.NumberNIR,
		"nir_date":       form.DateNIR,
		"invoice_number": form.NumberInvoice,
		"invoice_date":   form.DateInvoice,
		"provider":       form.Provider,
	})
}

func addProd(c *gin.Context) {
	nirNumber := c.Query("nir_number")
	productName := c.Query("product")
	amount := c.Query("amount")
	jsonPath := fmt.Sprintf("%s/indivitual_nir/%s.json", Username, nirNumber)

	fmt.Println("Adding product:", productName, "with amount:", amount, "to NIR number:", nirNumber)

	err, details := AddProduct(jsonPath, productName, amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"details": details,
		})
	}

}

func dataBaseInit(c *gin.Context) {
	// InitDB()
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default()) // Allows all origins, methods, and headers

	router.POST("/nir/create", nirCreate)
	router.GET("/nir/add", addProd)
	router.GET("/nir/initdb", dataBaseInit)

	dbutils.InitDB()
	dbutils.TransferDB()

	router.Run(":8082")
}
