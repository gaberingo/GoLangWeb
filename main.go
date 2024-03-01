package main

import (
	"log"
	"net/http"
	"ringo_web/config"
	"ringo_web/controllers/controller_categ"
	"ringo_web/controllers/controller_home"
)

func main() {
	config.ConnectDB()

	// 1.Homepage
	http.HandleFunc("/", controller_home.Index)

	// 2. Categories
	http.HandleFunc("/categories", controller_categ.Index)
	http.HandleFunc("/categories/add", controller_categ.Add)
	http.HandleFunc("/categories/edit", controller_categ.Edit)
	http.HandleFunc("/categories/delete", controller_categ.Delete)

	log.Println("Server running on port 8000")

	http.ListenAndServe(":8000", nil)
}
