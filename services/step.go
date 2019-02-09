package services

import (
	"recipe-book/db"
	"recipe-book/models"
)

func CreateStep(s *models.Step) error {
	if err := db.Connection.Create(&s).Error; err != nil {
		return err
	}

	return nil
}

func DeleteStep(id int) error {
	var step models.Step

	if err := db.Connection.Delete(&step, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateStep(s *models.Step, id int) (models.Step, error) {
	var step models.Step

	db.Connection.Find(&step, id)
	step.Number = s.Number
	step.Content = s.Content
	if err := db.Connection.Save(&step).Error; err != nil {
		return models.Step{}, err
	}

	return step, nil
}
