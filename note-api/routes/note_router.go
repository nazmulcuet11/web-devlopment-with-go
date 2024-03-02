package routes

import "github.com/gorilla/mux"

func ConfigureNoteRoutes(r *mux.Router) {
	ConfigureNoteApiRoutes(r)
	ConfigureNoteTemplateRoutes(r)
}
