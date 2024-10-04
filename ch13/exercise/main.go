package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func IpLoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		options := &slog.HandlerOptions{Level: slog.LevelInfo}
		handler := slog.NewJSONHandler(os.Stdout, options)
		mySlog := slog.New(handler)
		mySlog.Info("IP logger", "ip", r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

type CurrentTime struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func NewCurrentTimeFromTime(t time.Time) CurrentTime {
	return CurrentTime{
		DayOfWeek:  t.Weekday().String(),
		DayOfMonth: t.Day(),
		Month:      int(t.Month()),
		Year:       t.Year(),
		Hour:       t.Hour(),
		Minute:     t.Minute(),
		Second:     t.Second(),
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		if strings.Contains(r.Header.Get("Accept"), "application/json") {
			currentTime := NewCurrentTimeFromTime(now)
			timeAsJSON, err := json.Marshal(&currentTime)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("error getting latest time"))
				return
			}
			w.Write(timeAsJSON)
			return
		}
		formattedTime := now.Format(time.RFC3339)
		w.Write([]byte(formattedTime + "\n"))
	})

	wrappedMux := IpLoggingMiddleware(mux)

	s := http.Server{
		Addr:    ":9090",
		Handler: wrappedMux,
	}

	s.ListenAndServe()
}
