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
that a request is allowed to run.
It should return a func(http.Handler) http.Handler.
*/

func MiddleWare(handler http.Handler, milliseconds int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, time.Duration(milliseconds)*time.Millisecond)
		defer cancel()
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result, err := Process(r.Context())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(result))
}

func Process(ctx context.Context) (string, error) {
	wait := rand.Intn(4000)
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return "Hello!", nil
	case <-ctx.Done():
		return "Too slow!", ctx.Err()
	}
}

func main() {

	mtx := http.NewServeMux()
	helloHandler := HelloHandler{}
	wrappedHandler := MiddleWare(helloHandler, 2000)
	mtx.Handle("/hello", wrappedHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mtx,
	}

	server.ListenAndServe()
}
