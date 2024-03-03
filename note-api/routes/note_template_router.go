package routes

import (
	"net/http"
	"note-api/controllers"

	"github.com/gorilla/mux"
)

type NoteTemplateRouter struct {
	r          *mux.Router
	controller *controllers.NoteTemplatleController
}

func NewNoteTemplateRouter(
	r *mux.Router,
	controller *controllers.NoteTemplatleController,
) *NoteTemplateRouter {
	return &NoteTemplateRouter{r, controller}
}

func (router *NoteTemplateRouter) ConfigureRoutes() {
	c := router.controller
	r := router.r
	r.HandleFunc("/", c.GetNotes)
	r.HandleFunc("/notes/add", c.AddNote)
	r.HandleFunc("/notes/save", c.SaveNote)
	r.HandleFunc("/notes/edit/{id}", c.EditNote)
	r.HandleFunc("/notes/update/{id}", c.UpdateNote)
	r.HandleFunc("/notes/delete/{id}", c.DeleteNote)
}

func (router *NoteTemplateRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.r.ServeHTTP(w, r)
}
