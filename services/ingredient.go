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

func CreateIngredient(i *models.Ingredient) ([]models.Ingredient, error) {
	if err := db.Connection.Create(i).Error; err != nil {
		return nil, err
	}

	var ingredients []models.Ingredient
	if err := db.Connection.Find(&ingredients).Error; err != nil {
		return nil, err
	}

	return ingredients, nil
}

func DeleteIngredient(id int) ([]models.Ingredient, error) {
	var ingredient models.Ingredient

	if err := db.Connection.Delete(&ingredient, id).Error; err != nil {
		return nil, err
	}

	var ingredients []models.Ingredient
	if err := db.Connection.Find(&ingredients).Error; err != nil {
		return nil, err
	}

	return ingredients, nil
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

func BuildIngredientDTOList(ingredients []models.Ingredient) []models.IngredientDTO {
	var ingredientDTOList []models.IngredientDTO
	for _, v := range ingredients {
		ingredientDTOList = append(ingredientDTOList, BuildIngredientDTO(v))
	}

	return ingredientDTOList
}

func BuildIngredientDTO(ingredient models.Ingredient) models.IngredientDTO {
	ingredientDTO := models.IngredientDTO{
		ID:        ingredient.ID,
		CreatedAt: ingredient.CreatedAt,
		UpdatedAt: ingredient.UpdatedAt,
		DeletedAt: ingredient.DeletedAt,
		Name:      ingredient.Name,
		Notes:     ingredient.Notes,
		Recipes:   ingredient.Recipes}

	return ingredientDTO
}
