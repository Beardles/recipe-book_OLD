package models

import (
	"github.com/jinzhu/gorm"
)

type Ingredient struct {
	gorm.Model
	Name    string
	Recipes []*Recipe `gorm:"many2many:recipe_ingredients;"`
}
