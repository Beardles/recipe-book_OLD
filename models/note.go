package models

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	RecipeID int
	Content  string
}
