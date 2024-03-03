package repositories

import (
	"note-api/models"
)

type NoteRepository interface {
	GetAllNotes() []models.Note
	GetNoteById(id string) (models.Note, error)
	AddNote(note models.Note) (models.Note, error)
	AddNotes(notes []models.Note) error
	UpdateNote(newNote models.Note) error
	DeleteNoteById(id string) error
	DeleteNote(note models.Note) error
}
