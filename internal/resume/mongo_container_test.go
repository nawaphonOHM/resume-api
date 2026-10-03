//go:build testcontainers

// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"
	"testing"
	"time"

	"github.com/nawaphonOHM/whatever/pkg/mongodb"
	tcmongodb "github.com/nawaphonOHM/whatever/pkg/testcontainers/mongodb"
)

var (
	sharedContainer *tcmongodb.Container
	sharedClient    *mongodb.Client
	initErr         error
)

func closeSharedContainer() {
	if sharedContainer == nil {
		return
	}
	if err := sharedContainer.Terminate(context.Background()); err != nil {
		return
	}
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	sharedContainer, initErr = tcmongodb.Run(ctx, tcmongodb.WithDatabase("resume_test"))
	if initErr == nil {
		sharedClient, initErr = sharedContainer.Client(ctx)
	}
	defer closeSharedContainer()
	m.Run()
}

type testMongoHelper struct {
	t      *testing.T
	client *mongodb.Client
	repo   *MongoRepository
}

func newTestMongoHelper(t *testing.T) *testMongoHelper {
	t.Helper()
	if initErr != nil {
		t.Fatalf("failed to initialize mongodb testcontainer: %v", initErr)
	}
	ctx := context.Background()
	if err := sharedClient.Database().Drop(ctx); err != nil {
		t.Fatalf("failed to clean database: %v", err)
	}
	return &testMongoHelper{
		t:      t,
		client: sharedClient,
		repo:   NewMongoRepository(sharedClient),
	}
}

func (h *testMongoHelper) insert(col string, ver int32, val any) {
	h.t.Helper()
	coll := h.client.Collection(col)
	if coll == nil {
		h.t.Fatal("collection is nil")
	}
	doc := Document[any]{
		CreateDate: time.Now().UTC(),
		Version:    ver,
		Value:      val,
	}
	if _, err := coll.InsertOne(context.Background(), doc); err != nil {
		h.t.Fatalf("failed to insert doc: %v", err)
	}
}
