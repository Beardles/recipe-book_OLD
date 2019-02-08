package db

import (
	"recipe-book/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Connection *gorm.DB
var err error

func Open() {
	Connection, err = gorm.Open("postgres", "host=web553.webfaction.com port=5432 user=gorm dbname=recipe_book_dev password=Jb294784@")
	if err != nil {
		panic(err)
	}

}

func Bootstrap() {
	Connection.AutoMigrate(&models.Ingredient{}, &models.Note{}, &models.Step{}, &models.Recipe{})
}

func Close() {
	Connection.Close()
}
