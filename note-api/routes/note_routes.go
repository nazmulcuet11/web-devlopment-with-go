package routes

import (
	"note-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigureNoteApiRoutes(c *controllers.NoteApiController, r *mux.Router) {
	r.HandleFunc("/api/notes", c.GetNote).Methods("GET")
	r.HandleFunc("/api/notes", c.PostNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", c.PutNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", c.DeleteNote).Methods("DELETE")
}

func ConfigureNoteTemplateRoutes(c *controllers.NoteTemplatleController, r *mux.Router) {
	r.HandleFunc("/", c.GetNotes)
	r.HandleFunc("/notes/add", c.AddNote)
	r.HandleFunc("/notes/save", c.SaveNote)
	r.HandleFunc("/notes/edit/{id}", c.EditNote)
	r.HandleFunc("/notes/update/{id}", c.UpdateNote)
	r.HandleFunc("/notes/delete/{id}", c.DeleteNote)
}
