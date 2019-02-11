package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Ingredient struct {
	gorm.Model
	Name    string
	Recipes []*Recipe `gorm:"many2many:recipe_ingredients;"`
}

type IngredientDTO struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name"`
	Recipes   []*Recipe  `json:"recipes"`
}
