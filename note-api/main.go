package main

import (
	"fmt"
	"net/http"
	"note-api/controllers"
	"note-api/repositories"
	"note-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// populate test data
	repo := repositories.NewInMemoryNoteRepository()
	repo.PopulateTestNotes()

	apiController := controllers.NewNoteApiController(repo)
	templateController := controllers.NewNoteTemplatleController(repo)
	r := mux.NewRouter().StrictSlash(false)
	routes.ConfigureNoteApiRoutes(apiController, r)
	routes.ConfigureNoteTemplateRoutes(templateController, r)

	fmt.Printf("Listening at: 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
