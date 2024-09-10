package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler( w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	greeting := fmt.Sprintf("Hello, %s!", name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(greeting))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/greet/{name}", greetHandler).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}