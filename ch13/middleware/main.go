package main

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		dur := time.Since(start)
		slog.Info("request time", "path", r.URL.Path, "duration", dur)
	})
}

var securityMsg = []byte("you did not give the secret password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func main() {

	// configure terrible security providers
	terribleSecurity := TerribleSecurityProvider("secret")

	// instantiate mux
	mux := http.NewServeMux()

	// register the handler
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!\n"))
	})

	// wrap the mux with middleware
	wrappedMux := terribleSecurity(RequestTimer(mux))

	// instatiate the server
	s := http.Server{
		Addr:    ":8080",
		Handler: wrappedMux,
	}

	// start the server
	s.ListenAndServe()
}
