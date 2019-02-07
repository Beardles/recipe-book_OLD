package main

import (
	"recipe-book/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("postgres", "host=web553.webfaction.com port=5432 user=gorm dbname=recipe_book_dev password=Jb294784@")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Ingredient{}, &models.Note{}, &models.Step{}, &models.Recipe{})
	defer db.Close()
}
