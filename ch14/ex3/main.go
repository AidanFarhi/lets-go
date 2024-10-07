package main

import (
	"context"
	"fmt"
)

type Level string

const Debug Level = "debug"
const Info Level = "info"

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	// TODO get a logging level out of the context and assign it to inLevel
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func SetLogLevel(ctx context.Context, logLevel Level) context.Context {
	return context.WithValue(ctx, "logLevel", logLevel)
}

func GetLogLevel(ctx context.Context) Level {
	return ctx.Value("loglevel").(Level)
}

func main() {

}
