package repositories

import (
	"errors"
	"log"
	"note-api/models"
	"strconv"
	"time"
)

var notesMap = make(map[string]models.Note)
var id int = 0

func GetAllNotes() []models.Note {
	notes := make([]models.Note, len(notesMap))
	i := 0
	for _, note := range notesMap {
		notes[i] = note
		i++
	}
	return notes
}

func GetNoteById(id string) (models.Note, error) {
	note, ok := notesMap[id]
	if !ok {
		return models.Note{}, errors.New("note does not exist")
	}
	return note, nil
}

func AddNote(note models.Note) (models.Note, error) {
	_, ok := notesMap[note.Id]
	if ok {
		return models.Note{}, errors.New("note exists")
	}
	id++
	note.Id = strconv.Itoa(id)
	note.CreatedOn = time.Now()
	notesMap[note.Id] = note
	log.Println("Added note: ", note)
	return note, nil
}

func AddNotes(notes []models.Note) error {
	for _, note := range notes {
		_, err := AddNote(note)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateNote(newNote models.Note) error {
	existingNote, ok := notesMap[newNote.Id]
	if !ok {
		return errors.New("note does not exist")
	}
	newNote.CreatedOn = existingNote.CreatedOn
	notesMap[newNote.Id] = newNote
	return nil
}

func DeleteNoteById(id string) error {
	_, ok := notesMap[id]
	if !ok {
		return errors.New("note does not exist")
	}
	delete(notesMap, id)
	return nil
}

func DeleteNote(note models.Note) error {
	return DeleteNoteById(note.Id)
}
