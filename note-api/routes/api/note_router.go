package api

import (
	controllers "note-api/controllers/api"

	"github.com/gorilla/mux"
)

func ConfigureNoteApiRoutes(r *mux.Router) {
	r.HandleFunc("/api/notes", controllers.GetNote).Methods("GET")
	r.HandleFunc("/api/notes", controllers.PostNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", controllers.PutNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", controllers.DeleteNote).Methods("DELETE")
}
