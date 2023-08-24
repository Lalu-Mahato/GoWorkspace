package services

import (
	"notes_app/models"
	"notes_app/repositories"
)

type NoteService struct {
	noteRepository *repositories.NoteRepository
}

func NewNoteService(noteRepository *repositories.NoteRepository) *NoteService {
	return &NoteService{noteRepository: noteRepository}
}

func (noteService *NoteService) CreateNote(note *models.Note) error {
	err := noteService.noteRepository.Create(note)
	if err != nil {
		return err
	}
	return nil
}

func (noteService *NoteService) FindNotes() ([]models.Note, error) {
	notes, err := noteService.noteRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return notes, nil
}
