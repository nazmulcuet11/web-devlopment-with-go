package main

import (
	"fmt"
	"net/http"
	"note-api/controllers"
	"note-api/routes"
	"note-api/test_data"

	"github.com/gorilla/mux"
)

func main() {
	// populate test data
	test_data.PopulateTestNotes()

	apiController := controllers.NewNoteApiController()
	templateController := controllers.NewNoteTemplatleController()
	r := mux.NewRouter().StrictSlash(false)
	apiRouter := routes.NewNoteApiRouter(r, apiController)
	templateRouter := routes.NewNoteTemplateRouter(r, templateController)
	noteRouter := routes.NewNoteRouter(r, apiRouter, templateRouter)
	noteRouter.ConfigureRoutes()

	fmt.Printf("Listening at: 8080")
	err := http.ListenAndServe(":8080", noteRouter)
	if err != nil {
		panic(err)
	}
}
