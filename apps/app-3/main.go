package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("App 3: Main function executed!")

	response, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var result struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Println("joke:", string(result.Value))

}
