package routes

import (
	"note-api/routes/api"
	"note-api/routes/template"

	"github.com/gorilla/mux"
)

func ConfigureNoteRoutes(r *mux.Router) {
	api.ConfigureNoteApiRoutes(r)
	template.ConfigureNoteTemplateRoutes(r)
}
