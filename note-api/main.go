package main

import (
	"fmt"
	"log"
	"net/http"
	"note-api/controllers"
	"note-api/middlewares"
	"note-api/repositories"
	"note-api/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
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

	handler := alice.New(
		middlewares.CustomLoggingHandler,
		middlewares.FileLoggingHandler,
		handlers.CompressHandler,
	)
	fmt.Println("Server started, listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", handler.Then(r)))
}
