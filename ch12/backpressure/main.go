package main

import (
	"errors"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func NewPressureGauge(limit int) *PressureGauge {
	return &PressureGauge{
		ch: make(chan struct{}, limit),
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func longProcess() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	pg := NewPressureGauge(10_000)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(longProcess()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many requests"))
		}
	})
	http.ListenAndServe(":7777", nil)
}
