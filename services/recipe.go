package services

import (
	"recipe-book/db"
	"recipe-book/models"
)

func GetRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe

	if err := db.Connection.Find(&recipes).Error; err != nil {
		return nil, err
	}

	return recipes, nil
}

func GetRecipe(id int) (models.Recipe, error) {
	var recipe models.Recipe
	var ingredients []models.Ingredient
	var notes []models.Note
	var steps []models.Step

	if err := db.Connection.Find(&recipe, id).Error; err != nil {
		return models.Recipe{}, err
	}

	if err := db.Connection.Model(&recipe).Related(&ingredients).Error; err != nil {
		return models.Recipe{}, err
	}

	if err := db.Connection.Model(&recipe).Related(&steps).Error; err != nil {
		return models.Recipe{}, err
	}

	if err := db.Connection.Model(&recipe).Related(&notes).Error; err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func CreateRecipe(r *models.Recipe) error {
	if err := db.Connection.Create(r).Error; err != nil {
		return err
	}

	return nil
}

func DeleteRecipe(id int) error {
	var recipe models.Recipe

	if err := db.Connection.Delete(&recipe, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateRecipe(r *models.Recipe, id int) (models.Recipe, error) {
	var recipe models.Recipe

	db.Connection.Find(&recipe, id)
	recipe.Title = r.Title
	recipe.Description = r.Description
	recipe.Rating = r.Rating
	if err := db.Connection.Save(&recipe).Error; err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}
