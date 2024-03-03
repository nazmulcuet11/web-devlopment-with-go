package main

import (
	"fmt"
	"log"
	"net/http"
	"note-api/controllers"
	"note-api/middlewares"
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

	fmt.Println("Server started, listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middlewares.LogginHandler(r)))
}
