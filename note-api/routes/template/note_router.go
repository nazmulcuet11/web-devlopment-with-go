package template

import (
	"note-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigureNoteTemplateRoutes(r *mux.Router) {
	c := controllers.NewNoteTemplatleController()
	r.HandleFunc("/", c.GetNotes)
	r.HandleFunc("/notes/add", c.AddNote)
	r.HandleFunc("/notes/save", c.SaveNote)
	r.HandleFunc("/notes/edit/{id}", c.EditNote)
	r.HandleFunc("/notes/update/{id}", c.UpdateNote)
	r.HandleFunc("/notes/delete/{id}", c.DeleteNote)
}
