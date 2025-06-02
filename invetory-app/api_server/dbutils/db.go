package dbutils

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var nirDb, _ = sql.Open("sqlite3", "../dbutils/nir.db")
var oldDb, _ = sql.Open("sqlite3", "../dbutils/produse.db")

func InitDB() {

	createNirTable := `
    CREATE TABLE IF NOT EXISTS nir (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        company TEXT,
        address TEXT,
        nir_number TEXT,
        nir_date TEXT,
        invoice_number TEXT,
        invoice_date TEXT,
        provider TEXT,
        delegate TEXT
    );`
	if _, err := nirDb.Exec(createNirTable); err != nil {
		log.Fatal(err)
	}

	// Add users table
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        company TEXT,
		address TEXT
    );`
	if _, err := nirDb.Exec(createUsersTable); err != nil {
		log.Fatal(err)
	}

	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT,
		product TEXT,
		TVA TEXT,
		unit TEXT,
		providerName TEXT,
		providerPrice TEXT,
		sellingPrice TEXT,
		productCount INTEGER
	);`
	if _, err := nirDb.Exec(createProductsTable); err != nil {
		log.Fatal(err)
	}
}

func getProductDetailsById(db *sql.DB, id int) (string, string, string, string, string, error) {
	row, err := db.Query("SELECT produs, TVA, u_m, pret_furnizor, pret_final FROM preturi WHERE id = ?", id)
	if err != nil {
		log.Fatal("Error querying product details:", err)
	}
	defer row.Close()

	var product, TVA, unit, providerPrice, sellingPrice string
	if row.Next() {
		err = row.Scan(&product, &TVA, &unit, &providerPrice, &sellingPrice)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
	} else {
		fmt.Println("No row found for id:", id)
	}

	return product, TVA, unit, providerPrice, sellingPrice, err
}

func addProductToDB(db *sql.DB, product, TVA, unit, providerPrice, sellingPrice string) error {
	user := "test"
	productCount := 0
	_, err := db.Exec("INSERT INTO products (user, product, TVA, unit, providerPrice, sellingPrice, productCount) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user, product, TVA, unit, providerPrice, sellingPrice, productCount)
	if err != nil {
		log.Fatal("Error inserting product into database:", err)
		return err
	}
	return nil
}

func TransferDB() {
	for i := 1; i < 756; i++ {
		// for i := 1; i < 4; i++ {
		product, TVA, unit, providerPrice, sellingPrice, _ := getProductDetailsById(oldDb, i)
		fmt.Printf("Product: %s, TVA: %s, Unit: %s, Provider Price: %s, Selling Price: %s\n",
			product, TVA, unit, providerPrice, sellingPrice)
		if product == "" || strings.Contains(product, "test") || len(product) < 7 {
			continue
		} else {
			if TVA == "0.009" || TVA == "0.09" {
				TVA = "9"
			} else if TVA == "0.019" || TVA == "UNDIFINED" || TVA == "0.19" {
				TVA = "19"
			}

			addProductToDB(nirDb, product, TVA, unit, providerPrice, sellingPrice)
		}
	}
}

func GetProductDetails(rowsPerPage int32, offset int32) ([]map[string]interface{}, int, error) {
	row, err := nirDb.Query("SELECT id, product, TVA, unit, providerPrice, sellingPrice, productCount FROM products")
	if err != nil {
		return nil, 0, err
	}

	var id, product, TVA, unit, providerPrice, sellingPrice, productCount string
	var productDetails []map[string]interface{}
	for row.Next() {
		err = row.Scan(&id, &product, &TVA, &unit, &providerPrice, &sellingPrice, &productCount)
		if err != nil {
			log.Fatal("Error scanning row:", err)
			return nil, 0, err
		}
		productDetails = append(productDetails, map[string]interface{}{
			"id":            id,
			"product":       product,
			"TVA":           TVA,
			"unit":          unit,
			"providerPrice": providerPrice,
			"sellingPrice":  sellingPrice,
			"productCount":  productCount,
		})
	}

	sort.Slice(productDetails, func(i, j int) bool {
		return productDetails[i]["product"].(string) < productDetails[j]["product"].(string)
	})

	return productDetails[offset : rowsPerPage+offset], len(productDetails), err
}

func GetProductDetailsByProductName(productName string) ([]map[string]interface{}, error) {
	row, err := nirDb.Query("SELECT id, product, TVA, unit, providerPrice, sellingPrice, productCount FROM products WHERE product LIKE ?", productName)
	if err != nil {
		return nil, err
	}
	var id, product, TVA, unit, providerPrice, sellingPrice, productCount string
	var productDetails []map[string]interface{}
	for row.Next() {
		err = row.Scan(&id, &product, &TVA, &unit, &providerPrice, &sellingPrice, &productCount)
		if err != nil {
			log.Fatal("Error scanning row:", err)
			return nil, err
		}
		productDetails = append(productDetails, map[string]interface{}{
			"id":            id,
			"product":       product,
			"TVA":           TVA,
			"unit":          unit,
			"providerPrice": providerPrice,
			"sellingPrice":  sellingPrice,
			"productCount":  productCount,
		})
	}
	return productDetails, nil
}
