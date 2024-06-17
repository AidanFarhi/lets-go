package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

var options = &slog.HandlerOptions{Level: slog.LevelDebug}
var handler = slog.NewJSONHandler(os.Stderr, options)
var jsonSlog = slog.New(handler)

func IPLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		jsonSlog.Info("ip found", "ip", ip)
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yolo"))
	})
	wrappedMux := IPLogger(mux)
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 120 * time.Second,
		Handler:      wrappedMux,
	}
	server.ListenAndServe()
}
