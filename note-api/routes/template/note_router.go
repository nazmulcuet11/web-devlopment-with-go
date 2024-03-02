package template

import (
	controllers "note-api/controllers/template"

	"github.com/gorilla/mux"
)

func ConfigureNoteTemplateRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.GetNotes)
	r.HandleFunc("/notes/add", controllers.AddNote)
	r.HandleFunc("/notes/save", controllers.SaveNote)
	r.HandleFunc("/notes/edit/{id}", controllers.EditNote)
	r.HandleFunc("/notes/update/{id}", controllers.UpdateNote)
	r.HandleFunc("/notes/delete/{id}", controllers.DeleteNote)
}
