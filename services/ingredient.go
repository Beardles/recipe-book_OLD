package services

import (
	"recipe-book/db"
	"recipe-book/models"
)

func GetIngredients() ([]models.Ingredient, error) {
	var ingredients []models.Ingredient

	if err := db.Connection.Find(&ingredients).Error; err != nil {
		return nil, err
	}

	return ingredients, nil
}

func GetIngredient(id int) (models.Ingredient, error) {
	var ingredient models.Ingredient

	if err := db.Connection.Find(&ingredient, id).Error; err != nil {
		return models.Ingredient{}, err
	}

	return ingredient, nil
}

func CreateIngredient(i *models.Ingredient) error {
	if err := db.Connection.Create(i).Error; err != nil {
		return err
	}

	return nil
}

func DeleteIngredient(id int) error {
	var ingredient models.Ingredient

	if err := db.Connection.Delete(&ingredient, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateIngredient(i *models.Ingredient, id int) (models.Ingredient, error) {
	var ingredient models.Ingredient

	db.Connection.Find(&ingredient, id)
	ingredient.Name = i.Name
	if err := db.Connection.Save(&ingredient).Error; err != nil {
		return models.Ingredient{}, err
	}

	return ingredient, nil
}
