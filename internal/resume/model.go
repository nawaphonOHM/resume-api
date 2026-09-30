// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

// Logo represents image source metadata for institutions and companies.
type Logo struct {
	Src     string `bson:"src" json:"src"`
	Surface string `bson:"surface,omitempty" json:"surface,omitempty"`
	Height  int32  `bson:"height,omitempty" json:"height,omitempty"`
	Width   int32  `bson:"width,omitempty" json:"width,omitempty"`
}

// SeniorProject represents senior project details.
type SeniorProject struct {
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
}

// Details represents personal contact and biographical info.
type Details struct {
	BirthDate   string `bson:"birthDate" json:"birthDate"`
	Email       string `bson:"email" json:"email"`
	Location    string `bson:"location" json:"location"`
	Nationality string `bson:"nationality" json:"nationality"`
	PhoneLabel  string `bson:"phoneLabel" json:"phoneLabel"`
}
