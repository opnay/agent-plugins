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
		envToken: os.Getenv("TYPESAFE_API_KEY"),
	}
	code := app.run(ctx, os.Args[1:])
	cancel()
	os.Exit(code)
}
