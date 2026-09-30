// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"context"

	"github.com/nawaphonOHM/whatever/pkg/mongodb"
)

type cursor interface {
	Next(context.Context) bool
	Decode(any) error
	Err() error
	Close(context.Context) error
}

func closeCursor(ctx context.Context, cur cursor) {
	if err := cur.Close(ctx); err != nil {
		return
	}
}

func advanceCursor(ctx context.Context, cur cursor) error {
	if !cur.Next(ctx) {
		if err := cur.Err(); err != nil {
			return err
		}
		return ErrNotFound
	}
	return nil
}

func fetchCursor(ctx context.Context, client *mongodb.Client, name string) (cursor, error) {
	if client == nil {
		return nil, mongodb.ErrNilClient
	}
	coll := client.Collection(name)
	if coll == nil {
		return nil, mongodb.ErrNilClient
	}
	pipeline := []map[string]any{
		{"$sort": map[string]int{"version": -1}},
		{"$limit": 1},
	}
	return coll.Aggregate(ctx, pipeline)
}

func decodeCursor[T any](ctx context.Context, cur cursor) (*T, error) {
	defer closeCursor(ctx, cur)
	if err := advanceCursor(ctx, cur); err != nil {
		return nil, err
	}
	var doc Document[T]
	if err := cur.Decode(&doc); err != nil {
		return nil, err
	}
	return &doc.Value, nil
}

func queryLatest[T any](ctx context.Context, client *mongodb.Client, name string) (*T, error) {
	cur, err := fetchCursor(ctx, client, name)
	if err != nil {
		return nil, err
	}
	return decodeCursor[T](ctx, cur)
}
