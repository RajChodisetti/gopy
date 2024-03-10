package main

import (
	"fmt"
	"gopy/api"
	"net/http"
	"os"
)

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	// Get the directory path from the request
	directory := r.URL.Query().Get("directory")

	// Call the API function with the directory path
	if directory == "" {
		wd, err := os.Getwd()
		if err == nil {
			fmt.Println(wd)
			api.Api(wd)
		}

	} else {
		fmt.Println("Directory is :}", directory)
		api.Api(directory)
	}

	// Write a response to the client
	fmt.Fprintf(w, "API request processed successfully")
}

func main() {
	// Register the handleAPIRequest function as the handler for the /api endpoint
	http.HandleFunc("/api", handleAPIRequest)

	http.ListenAndServe(":8080", nil)
}
