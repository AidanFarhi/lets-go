package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("only GET allowed"))
			return
		}
		w.Write([]byte(time.Now().Format(time.RFC3339)))
	})
	http.ListenAndServe(":8080", mux)
}
