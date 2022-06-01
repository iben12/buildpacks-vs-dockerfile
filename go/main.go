package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("New %s request on path \"%s\" at %s\n", r.Method, r.URL, time.Now())
	io.WriteString(w, "Hello, Docker Meetup!\n")
	fmt.Printf("%s request on path \"%s\" served at %s\n", r.Method, r.URL, time.Now())
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

// Initiate web server
func main() {
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
