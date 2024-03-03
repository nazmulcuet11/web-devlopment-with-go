package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"note-api/models"
	"note-api/repositories"

	"github.com/gorilla/mux"
)

type NoteApiController struct {
	repo repositories.NoteRepository
}

func NewNoteApiController(repo repositories.NoteRepository) *NoteApiController {
	return &NoteApiController{repo}
}

// HTTP Get - /api/notes
func (c *NoteApiController) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	notes := c.repo.GetAllNotes()
	j, err := json.Marshal(notes)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP Post - /api/notes
func (c *NoteApiController) PostNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	note, err = c.repo.AddNote(note)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	noteJson, err := json.Marshal(note)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(noteJson)
}

// HTTP Put - /api/notes/{id}
func (c *NoteApiController) PutNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newNote models.Note
	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	newNote.Id = vars["id"]
	err = c.repo.UpdateNote(newNote)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// HTTP Delete - /api/notes/{id}
func (c *NoteApiController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	err := c.repo.DeleteNoteById(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
