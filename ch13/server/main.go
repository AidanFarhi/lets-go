package main

import (
	"fmt"
	"net/http"
	"time"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}

func handleGoodbye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("goodbye\n"))
}

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write([]byte(fmt.Sprintf("hello %s\n", name)))
}

func handleCatGreeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("meow!\n"))
}

func handleDogGreeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bork!\n"))
}

func main() {

	catMux := http.NewServeMux()
	catMux.HandleFunc("GET /hello", handleCatGreeting)

	dogMux := http.NewServeMux()
	dogMux.HandleFunc("GET /hello", handleDogGreeting)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /goodbye", handleGoodbye)
	mux.HandleFunc("GET /hello/{name}", handleGreeting)
	mux.Handle("GET /hello", HelloHandler{})
	mux.Handle("/cat/", http.StripPrefix("/cat", dogMux))
	mux.Handle("/dog/", http.StripPrefix("/dog", dogMux))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	s.ListenAndServe()
}
