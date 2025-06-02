package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetImageByStringSerpApi(productName string) (string, error) {
	apiKey := "e950fb8bf49678880f348cb9248ed66fe2899531996e3c66044a2f084747d415"
	url := fmt.Sprintf("https://serpapi.com/search.json?q=%s&tbm=isch&api_key=%s", productName, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if images, ok := result["images_results"].([]interface{}); ok && len(images) > 0 {
		if first, ok := images[0].(map[string]interface{}); ok {
			if link, ok := first["original"].(string); ok {
				return link, nil
			}
		}
	}
	return "", nil
}
