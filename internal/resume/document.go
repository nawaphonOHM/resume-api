// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

import (
	"time"
)

// Document wraps a MongoDB document payload with metadata fields.
type Document[T any] struct {
	Value      T         `bson:"value"`
	CreateDate time.Time `bson:"create_date"`
	Version    int32     `bson:"version"`
}
