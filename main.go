package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func coutingTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)

		log.Printf("Completed in %v", time.Since(start))
	})
}

func helloHandler( w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, exists := vars["name"]

	if !exists || name == "" {
		http.Error(w, "Name parameter is missing", http.StatusBadRequest)
		return
	}

	greeting := fmt.Sprintf("Hello, %s!", name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(greeting))
}

func main() {
	r := mux.NewRouter()

	r.Use(coutingTimeMiddleware)

	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/greet/{name:[a-zA-Z]*}", greetHandler).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}