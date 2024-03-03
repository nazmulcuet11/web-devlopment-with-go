package api

import (
	"note-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigureNoteApiRoutes(r *mux.Router) {
	c := controllers.NewNoteApiController()
	r.HandleFunc("/api/notes", c.GetNote).Methods("GET")
	r.HandleFunc("/api/notes", c.PostNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", c.PutNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", c.DeleteNote).Methods("DELETE")
}
