// Package main provides the entry point for the resume REST API server.
package main

import (
	"context"
	"os"

	"github.com/nawaphonOHM/resume-api/internal/resume"
	"github.com/nawaphonOHM/whatever/pkg/logging"
	"github.com/nawaphonOHM/whatever/pkg/mongodb"
	"github.com/nawaphonOHM/whatever/pkg/rest"
)

func run() error {
	ctx := context.Background()
	client, err := mongodb.Connect(ctx)
	if err != nil {
		return err
	}
	repo := resume.NewMongoRepository(client)
	service := resume.NewService(repo)
	handler := resume.NewHandler(service)
	bp := rest.NewBluePrint().AddAPIs(resume.Routes(handler))
	bp.Meta().Cors().WithAllowOrigin("resume.ohm-mho.space").WithAllowHTTPMethods(rest.GET)
	return rest.StartREST(bp)
}

func main() {
	if err := run(); err != nil {
		logging.Error("server stopped with error", "error", err)
		os.Exit(1)
	}
}
