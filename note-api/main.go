package main

import (
	"fmt"
	"net/http"
	"note-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	routes.ConfigureNoteRoutes(r)
	fmt.Printf("Listening at: 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
