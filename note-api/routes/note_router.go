package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type NoteRouter struct {
	r              *mux.Router
	apiRouter      *NoteApiRouter
	templateRouter *NoteTemplateRouter
}

func NewNoteRouter(
	r *mux.Router,
	apiRouter *NoteApiRouter,
	templateRouter *NoteTemplateRouter,
) *NoteRouter {
	return &NoteRouter{r, apiRouter, templateRouter}
}

func (router *NoteRouter) ConfigureRoutes() {
	router.apiRouter.ConfigureRoutes()
	router.templateRouter.ConfigureRoutes()
}

func (router *NoteRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.r.ServeHTTP(w, r)
}
