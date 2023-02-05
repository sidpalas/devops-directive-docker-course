package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get("http://localhost:" + port + "/ping")
	if err != nil {
		log.Fatal(err)
	}

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Argh! Broken")
		os.Exit(1)
	}
}
