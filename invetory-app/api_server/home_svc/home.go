package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func homeSummaryBoard(c *gin.Context) {
	response := gin.H{
		"providers": gin.H{
			"unit":  []string{"S.C. Facos S.A.", "S.C. Taviany S.R.L.", "S.C. Alfa S.R.L.", "S.C. Beta S.R.L."},
			"total": []int{40, 30, 20, 10},
		},
		"products": gin.H{
			"unit":  []string{"Franzela 300g", "Franzela 500g", "Kent 19", "Apa plata Aqua 2l"},
			"total": []int{100, 200, 150, 50},
		},
		"sales": gin.H{
			"unit":  []string{"Ianuare", "Februarie", "Martie", "Aprilie"},
			"total": []int{1000, 2000, 1500, 500},
		},
		"purchases": gin.H{
			"unit":  []string{"Ianuare", "Februarie", "Martie", "Aprilie"},
			"total": []int{800, 1800, 1200, 600},
		},
	}
	c.JSON(http.StatusOK, response)
}

func homeDetails(c *gin.Context) {
	responseDetails := []map[string]interface{}{
		{"id": 1, "numberNIR": "570", "dateNIR": "2023-10-01", "numberInvoice": "12345",
			"dateInvoice": "2023-10-02", "provider": "S.C. Facos S.A.", "total": 100},
		{"id": 2, "numberNIR": "571", "dateNIR": "2023-10-03", "numberInvoice": "12346",
			"dateInvoice": "2023-10-04", "provider": "S.C. Taviany S.R.L.", "total": 200},
		{"id": 3, "numberNIR": "572", "dateNIR": "2023-10-05", "numberInvoice": "12347",
			"dateInvoice": "2023-10-06", "provider": "S.C. Alfa S.R.L.", "total": 150},
		{"id": 4, "numberNIR": "573", "dateNIR": "2023-10-07", "numberInvoice": "12348",
			"dateInvoice": "2023-10-08", "provider": "S.C. Beta S.R.L.", "total": 50},
		{"id": 5, "numberNIR": "574", "dateNIR": "2023-10-09", "numberInvoice": "12349",
			"dateInvoice": "2023-10-10", "provider": "S.C. Facos S.A.", "total": 300},
		{"id": 6, "numberNIR": "575", "dateNIR": "2023-10-11", "numberInvoice": "12350",
			"dateInvoice": "2023-10-12", "provider": "S.C. Taviany S.R.L.", "total": 400},

		{"id": "", "numberNIR": "", "dateNIR": "", "numberInvoice": "",
			"dateInvoice": "", "provider": "", "total": ""},
	}
	c.JSON(http.StatusOK, responseDetails)
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default()) // Allows all origins, methods, and headers

	router.GET("home/summary", homeSummaryBoard)
	router.GET("home/details", homeDetails)

	router.Run(":8081")
}
