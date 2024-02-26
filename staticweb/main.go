package main

import (
	"fmt"
	"log"
	"net/http"
)

//
//type messageHandler struct {
//	message string
//}
//
//func (h messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	fmt.Sprint(w, h.message)
//}

//func messageHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "some message")
//}

func messageHandler(message string) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, message)
	}
	return http.HandlerFunc(handler)
}

func main() {
	//mux := http.NewServeMux()

	//fs := http.FileServer(http.Dir("public"))
	//mux.Handle("/", fs)

	//mh1 := messageHandler{"Welcome to go web development"}
	//mux.Handle("/welcome", &mh1)
	//
	//mh2 := messageHandler{"net/http is awesome"}
	//mux.Handle("/message", &mh2)

	//mux.Handle("/message", http.HandlerFunc(messageHandler))

	//mux.Handle("/welcome", messageHandler("Welcome to go web development"))
	//mux.Handle("/message", messageHandler("net/http is awesome"))
	//
	//log.Println("Listening at: 8080")
	//err := http.ListenAndServe(":8080", mux)

	http.Handle("/welcome", messageHandler("Welcome to go web development"))
	http.Handle("/message", messageHandler("net/http is awesome"))

	log.Println("Listening at: 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
