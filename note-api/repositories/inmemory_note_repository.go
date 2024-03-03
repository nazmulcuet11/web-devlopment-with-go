package repositories

import (
	"errors"
	"log"
	"note-api/models"
	"note-api/test_data"
	"sort"
	"strconv"
	"time"
)

type InMemoryNoteRepository struct {
	id       int
	notesMap map[string]models.Note
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
	return &InMemoryNoteRepository{
		id:       0,
		notesMap: make(map[string]models.Note),
	}
}

func (r *InMemoryNoteRepository) GetAllNotes() []models.Note {
	notes := make([]models.Note, len(r.notesMap))
	i := 0
	for _, note := range r.notesMap {
		notes[i] = note
		i++
	}
	sort.Slice(notes, func(x, y int) bool {
		return notes[x].Id <= notes[y].Id
	})
	return notes
}

func (r *InMemoryNoteRepository) GetNoteById(id string) (models.Note, error) {
	note, ok := r.notesMap[id]
	if !ok {
		return models.Note{}, errors.New("note does not exist")
	}
	return note, nil
}

func (r *InMemoryNoteRepository) AddNote(note models.Note) (models.Note, error) {
	_, ok := r.notesMap[note.Id]
	if ok {
		return models.Note{}, errors.New("note exists")
	}
	r.id++
	note.Id = strconv.Itoa(r.id)
	note.CreatedOn = time.Now()
	r.notesMap[note.Id] = note
	log.Println("Added note: ", note)
	return note, nil
}

func (r *InMemoryNoteRepository) AddNotes(notes []models.Note) error {
	for _, note := range notes {
		_, err := r.AddNote(note)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *InMemoryNoteRepository) UpdateNote(newNote models.Note) error {
	existingNote, ok := r.notesMap[newNote.Id]
	if !ok {
		return errors.New("note does not exist")
	}
	newNote.CreatedOn = existingNote.CreatedOn
	r.notesMap[newNote.Id] = newNote
	log.Println("Updated note: ", newNote)
	return nil
}

func (r *InMemoryNoteRepository) DeleteNoteById(id string) error {
	_, ok := r.notesMap[id]
	if !ok {
		return errors.New("note does not exist")
	}
	delete(r.notesMap, id)
	log.Println("Deleted note with id: ", id)
	return nil
}

func (r *InMemoryNoteRepository) DeleteNote(note models.Note) error {
	return r.DeleteNoteById(note.Id)
}

func (r *InMemoryNoteRepository) PopulateTestNotes() {
	r.AddNotes(test_data.TestNotes)
}
