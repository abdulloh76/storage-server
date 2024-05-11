package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/abdulloh76/storage-server/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	http.HandleFunc("/upload", handlers.HandleUpload)
	http.HandleFunc("/download/", handlers.HandleDownload)

	fmt.Println("Server running on http://localhost:" + PORT)
	err = http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
