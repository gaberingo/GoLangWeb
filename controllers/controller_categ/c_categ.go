package controller_categ

import (
	"fmt"
	"html/template"
	"net/http"
	entities "ringo_web/entities"
	models_categ "ringo_web/models/models_categ"
	"time"
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
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()
		fmt.Println(category)

		ok := models_categ.Create(category)
		if !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
}

func Delete(w http.ResponseWriter, r *http.Request) {
}
