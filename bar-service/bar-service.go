package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()
	log.Println("Starting http listener on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		version := os.Getenv("SERVICE_VERSION")
		fmt.Fprintln(w, "hello from bar-service:", version)
	})
	return router
}
