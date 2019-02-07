package models

import (
	"github.com/jinzhu/gorm"
)

type Recipe struct {
	gorm.Model
	Title       string
	Description string
	Rating      int
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;"`
	Steps       []Step
	Notes       []Note
}
