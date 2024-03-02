package routes

import (
	"note-api/controllers"

	"github.com/gorilla/mux"
)

func ConfigureNoteTemplateRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.GetNotesTemplateHandler)
	r.HandleFunc("/notes/add", controllers.AddNoteTemplateHandler)
	r.HandleFunc("/notes/save", controllers.SaveNoteTemplateHandler)
	r.HandleFunc("/notes/edit/{id}", controllers.EditNoteTemplateHandler)
	r.HandleFunc("/notes/update/{id}", controllers.UpdateNoteTemplateHandler)
	r.HandleFunc("/notes/delete/{id}", controllers.DeleteNoteTemplateHandler)
}
