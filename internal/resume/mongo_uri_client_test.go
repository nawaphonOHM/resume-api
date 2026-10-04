//go:build realmongo

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"os"
	"testing"

	"github.com/nawaphonOHM/whatever/v2/pkg/testing/mongodb"
)

const mongoURIEnv = "MONGODB_URI"

func requireMongoURI(t *testing.T) string {
	t.Helper()
	rawURI := os.Getenv(mongoURIEnv)
	if rawURI == "" {
		t.Skip("skipping MongoDB integration test: MONGODB_URI not set")
	}
	return rawURI
}

func connectURIRepository(t *testing.T) *MongoRepository {
	t.Helper()
	client := mongodb.NewTestClientURI(t, requireMongoURI(t))
	return NewMongoRepository(client.Client())
}
