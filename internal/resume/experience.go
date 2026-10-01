// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

// ClientItem represents client details within a work experience.
type ClientItem struct {
	Logo *Logo  `bson:"logo,omitempty" json:"logo,omitempty"`
	Name string `bson:"name" json:"name"`
}

// ExperienceItem represents a work experience entry.
type ExperienceItem struct {
	Company         string      `bson:"company" json:"company"`
	Role            string      `bson:"role" json:"role"`
	Period          string      `bson:"period" json:"period"`
	Location        string      `bson:"location" json:"location"`
	CompanyLogo     *Logo       `bson:"companyLogo,omitempty" json:"companyLogo,omitempty"`
	Client          *ClientItem `bson:"client,omitempty" json:"client,omitempty"`
	EmploymentTypes []string    `bson:"employmentTypes" json:"employmentTypes"`
	Highlights      []string    `bson:"highlights" json:"highlights"`
	Technologies    []string    `bson:"technologies" json:"technologies"`
}
