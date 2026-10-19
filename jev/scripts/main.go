package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	app := application{
		in: os.Stdin, out: os.Stdout, errOut: os.Stderr,
		env: map[string]string{"TYPESAFE_API_KEY": os.Getenv("TYPESAFE_API_KEY"), "OPENROUTER_API_KEY": os.Getenv("OPENROUTER_API_KEY"), "OPENAI_API_KEY": os.Getenv("OPENAI_API_KEY")},
	}
	code := app.run(ctx, os.Args[1:])
	cancel()
	os.Exit(code)
}
