package main

import (
	"log"
	"os"

	"encoding/json"

	"fmt"
)

type ProductDetails struct {
	Product string `json:"product"`
	Amount  int32  `json:"amount"`
}

func CreateNirFile(nirDetails interface{}, nirNumber string) error {
	jsonPath := fmt.Sprintf("%s/indivitual_nir/%s.json", Username, nirNumber)
	details := []map[string]interface{}{
		{"nirDetails": nirDetails},
	}
	// Writes to the current working directory
	file, err := os.Create(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(details)
}

func AddProduct(jsonPath string, product string, amount string) (error, []map[string]interface{}) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	var details []map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode((&details)); err != nil {
		return err, nil
	}

	userProduct := map[string]interface{}{
		"product": map[string]interface{}{
			"product": product,
			"amount":  amount,
		},
	}

	details = append(details, userProduct)

	file, err = os.Create(jsonPath)
	if err != nil {
		return err, nil
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(details)

	return nil, details
}
