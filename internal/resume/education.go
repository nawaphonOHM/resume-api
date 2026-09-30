// Package resume defines domain data models, repository interfaces, and services for resume data.
package resume

// Education represents educational history and degree info.
type Education struct {
	SeniorProject   *SeniorProject `bson:"seniorProject,omitempty" json:"seniorProject,omitempty"`
	InstitutionLogo *Logo          `bson:"institutionLogo,omitempty" json:"institutionLogo,omitempty"`
	Institution     string         `bson:"institution" json:"institution"`
	Degree          string         `bson:"degree" json:"degree"`
	GPAX            string         `bson:"gpax" json:"gpax"`
	Period          string         `bson:"period" json:"period"`
}
