package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"note-api/models"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note models.Note

var noteStore = make(map[string]Note)
var id int = 0

// HTTP Get - /notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP Post - /notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// HTTP Put - /notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var noteToUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteToUpdate)
	if err != nil {
		panic(err)
	}
	vars := mux.Vars(r)
	k := vars["id"]
	if note, ok := noteStore[k]; ok {
		noteToUpdate.CreatedOn = note.CreatedOn
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
	} else {
		log.Printf("Could not find key of Note %s to delete\n", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

// HTTP Delete - /notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("Could not find key of Note %s to delete\n", k)
	}
	w.WriteHeader(http.StatusNoContent)
}
