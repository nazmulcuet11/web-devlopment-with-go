package routes

import (
	"net/http"
	"note-api/controllers"

	"github.com/gorilla/mux"
)

type NoteApiRouter struct {
	r          *mux.Router
	controller *controllers.NoteApiController
}

func NewNoteApiRouter(
	r *mux.Router,
	controller *controllers.NoteApiController,
) *NoteApiRouter {
	return &NoteApiRouter{r, controller}
}

func (router *NoteApiRouter) ConfigureRoutes() {
	c := router.controller
	r := router.r
	r.HandleFunc("/api/notes", c.GetNote).Methods("GET")
	r.HandleFunc("/api/notes", c.PostNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}", c.PutNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", c.DeleteNote).Methods("DELETE")
}

func (router *NoteApiRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.r.ServeHTTP(w, r)
}
