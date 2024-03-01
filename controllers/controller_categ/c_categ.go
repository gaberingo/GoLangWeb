package controller_categ

import (
	"html/template"
	"net/http"
	models_categ "ringo_web/models/models_categ"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := models_categ.GetAll()
	data := map[string]any{
		"categories": categories,
	}
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
}

func Edit(w http.ResponseWriter, r *http.Request) {
}

func Delete(w http.ResponseWriter, r *http.Request) {
}
