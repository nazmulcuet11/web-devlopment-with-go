package main

import (
	"fmt"
	"net/http"
	"note-api/routes"
	"note-api/test_data"

	"github.com/gorilla/mux"
)

func main() {
	// populate test data
	test_data.PopulateTestNotes()

	r := mux.NewRouter().StrictSlash(false)
	routes.ConfigureNoteRoutes(r)
	fmt.Printf("Listening at: 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
