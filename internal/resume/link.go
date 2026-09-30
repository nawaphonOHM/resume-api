// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

// LinkItem represents an external link entry.
type LinkItem struct {
	Logo  *Logo  `bson:"logo,omitempty" json:"logo,omitempty"`
	Label string `bson:"label" json:"label"`
	URL   string `bson:"url" json:"url"`
}
