package main

import (
	"context"
	"os"
	"os/signal"

	"opnay/jev/cli"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	app := cli.App{
		In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr,
		Env: map[string]string{"TYPESAFE_API_KEY": os.Getenv("TYPESAFE_API_KEY"), "OPENROUTER_API_KEY": os.Getenv("OPENROUTER_API_KEY"), "OPENAI_API_KEY": os.Getenv("OPENAI_API_KEY")},
	}
	code := app.Run(ctx, os.Args[1:])
	cancel()
	os.Exit(code)
}
