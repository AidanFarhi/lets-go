package main

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

/*
Create a middleware-generating function that creates a context with a timeout.
The function should have one parameter, which is the number of milliseconds
that a request is allowed to run. It should return a func(http.Handler) http.Handler.
*/

func NewTimeoutMiddleware(milliseconds int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancelFunc := context.WithTimeout(ctx, time.Millisecond*time.Duration(milliseconds))
			defer cancelFunc()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func process(ctx context.Context) (string, error) {
	wait := rand.Intn(3000)
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return "done\n", nil
	case <-ctx.Done():
		return "too slow\n", ctx.Err()
	}
}

func makeRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message, err := process(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(message))
}

func main() {

	timouteMiddleware := NewTimeoutMiddleware(2000)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api", makeRequest)
	wrappedMux := timouteMiddleware(mux)

	s := http.Server{
		Addr:    ":6060",
		Handler: wrappedMux,
	}

	s.ListenAndServe()
}
