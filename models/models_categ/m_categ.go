package models_categ

import (
	config "ringo_web/config"
	entities "ringo_web/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err.Error())
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
    INSERT INTO categories (name, created_at, updated_at)
    VALUES (?, ?, ?)
    `, category.Name, category.CreatedAt, category.UpdatedAt,
	)
	if err != nil {
		panic(err.Error())
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return lastInsertId > 0
}
