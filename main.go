// Package main provides the entry point for the resume REST API server.
package main

import (
	"context"
	"os"

	"github.com/nawaphonOHM/resume-api/internal/resume"
	"github.com/nawaphonOHM/whatever/v2/pkg/logging"
	"github.com/nawaphonOHM/whatever/v2/pkg/mongodb"
	"github.com/nawaphonOHM/whatever/v2/pkg/rest"
)

const allowedOrigin = "resume.ohm-mho.space"

func buildCorsSetting() *rest.CorsSetting {
	return rest.NewCorsSetting().WithAllowHTTPMethods(rest.GET).WithAllowOrigin(allowedOrigin)
}

func buildBlueprint(routes *rest.RRestAPIRegistration) *rest.BluePrint {
	return rest.NewBluePrint().WithAPIs(routes)
}

func initRoutes(client *mongodb.Client) *rest.RRestAPIRegistration {
	repo := resume.NewMongoRepository(client)
	service := resume.NewService(repo)
	handler := resume.NewHandler(service)
	return resume.Routes(handler)
}

func run() error {
	client, err := mongodb.Connect(context.Background())
	if err != nil {
		return err
	}
	return rest.StartREST(buildBlueprint(initRoutes(client)))
}

func main() {
	if err := run(); err != nil {
		logging.Error("server stopped with error", "error", err)
		os.Exit(1)
	}
}
