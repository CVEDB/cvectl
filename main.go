package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/cvedb/cvectl/pkg/cli"
)

func main() {
	if err := mainE(context.Background()); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}

func mainE(ctx context.Context) error {
	ctx, done := signal.NotifyContext(ctx, os.Interrupt)
	defer done()

	return cli.New().ExecuteContext(ctx)
}
