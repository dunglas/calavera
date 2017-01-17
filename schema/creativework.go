package schema

import "time"

// CreativeWork contains all data and metadata of a given page.
type CreativeWork struct {
	JsonLd
	Name         string     `json:"name"`
	Text         string     `json:"text"`
	About        string     `json:"about,omitempty"`
	Author       []Person   `json:"author,omitempty"`
	DateCreated  *time.Time `json:"dateCreated,omitempty"`
	DateModified *time.Time `json:"dateModified,omitempty"`
	InLanguage   string     `json:"inLanguage,omitempty"`
	License      string     `json:"license,omitempty"`
	Publisher    string     `json:"publisher,omitempty"`
}

// NewCreativeWork initializes a new CreativeWork instance with some sensitive default values.
func NewCreativeWork() *CreativeWork {
	return &CreativeWork{JsonLd: JsonLd{Context: "http://schema.org", Type: "CreativeWork"}}
}
