package routes

import (
	"note-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigureNoteRoutes(r *mux.Router) {
	r.HandleFunc("/api/notes", controllers.GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", controllers.PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", controllers.PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", controllers.DeleteNoteHandler).Methods("DELETE")
}
