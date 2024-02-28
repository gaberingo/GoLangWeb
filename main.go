package main

import (
	"log"
	"net/http"
	"ringo_web/config"
)

func main() {
	config.ConnectDB()
	log.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
