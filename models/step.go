package models

import (
	"github.com/jinzhu/gorm"
)

type Step struct {
	gorm.Model
	RecipeID int
	Number   int
	Content  string
}
