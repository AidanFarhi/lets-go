package main

import (
	"context"
	"fmt"
)

type Level string

const Debug Level = "debug"
const Info Level = "info"

type logLevelKey int

const key logLevelKey = 1

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	inLevel = GetLogLevel(ctx)
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func SetLogLevel(ctx context.Context, logLevel Level) context.Context {
	return context.WithValue(ctx, key, logLevel)
}

func GetLogLevel(ctx context.Context) Level {
	if ctx == nil {
		fmt.Println("ctx is nil")
	}
	return ctx.Value(key).(Level)
}

func main() {
	ctx := SetLogLevel(context.Background(), Debug)
	Log(ctx, Debug, "debug")
	Log(ctx, Info, "info")
	// ctx = SetLogLevel(ctx, Warn)
}
