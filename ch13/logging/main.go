package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {

	// basic logs
	slog.Info("message")
	slog.Warn("message")
	slog.Error("message")

	// custom log attributes
	userId := "fred"
	loginCount := 20
	slog.Info("user login", "id", userId, "login_count", loginCount)

	// JSON logging
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)
	lastLogin := time.Date(2023, 01, 01, 11, 50, 00, 00, time.UTC)
	mySlog.Debug("debug message", "id", userId, "last_login", lastLogin)
}
