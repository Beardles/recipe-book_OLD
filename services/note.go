package services

import (
	"recipe-book/db"
	"recipe-book/models"
)

func CreateNote(n *models.Note) error {
	if err := db.Connection.Create(&n).Error; err != nil {
		return err
	}

	return nil
}

func DeleteNote(id int) error {
	var note models.Note

	if err := db.Connection.Delete(&note, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateNote(n *models.Note, id int) (models.Note, error) {
	var note models.Note

	db.Connection.Find(&note, id)
	note.Content = n.Content
	if err := db.Connection.Save(&note).Error; err != nil {
		return models.Note{}, err
	}

	return note, nil
}
