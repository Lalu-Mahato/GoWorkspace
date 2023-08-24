package repositories

import (
	"notes_app/models"

	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (noteRepository *NoteRepository) Create(note *models.Note) error {
	err := noteRepository.db.Create(note).Error
	return err
}

func (noteRepository *NoteRepository) FindAll() ([]models.Note, error) {
	var notes []models.Note
	err := noteRepository.db.Find(&notes).Error
	return notes, err
}
