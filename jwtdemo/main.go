package main

import (
	"encoding/json"
	"fmt"
	"jwtdemo/auth"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Response struct {
	Text string `json:"text"`
}

type Token struct {
	Token string `json:"token"`
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error in request body")
		return
	}

	// simulate uservalidation
	if credentials.Username != "nazmul" && credentials.Password != "pass" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}

	tokenString, err := auth.GenerateToken(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error while signing", err)
		return
	}
	response := Token{tokenString}
	jsonResponse(response, w)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Missing Authorization header")
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	err := auth.VerifyToken(tokenString)
	if err != nil {
		response := Response{"Invalid token, error:" + err.Error()}
		jsonResponse(response, w)
	} else {
		response := Response{"Authorized to the system"}
		jsonResponse(response, w)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth", authHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
