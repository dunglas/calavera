package schema

type CreativeWork struct {
	JsonLd
	Name string `json:"name"`
	Text string `json:"text"`
	About string `json:"about,omitempty"`
	Author []Person `json:"author,omitempty"`
	DateCreated string `json:"dateCreated,omitempty"`
	DateModified string `json:"dateModified,omitempty"`
	InLanguage string `json:"inLanguage,omitempty"`
	License string `json:"license,omitempty"`
	Publisher string `json:"publisher,omitempty"`
}

func NewCreativeWork() *CreativeWork {
	return &CreativeWork{JsonLd: JsonLd{Context: "http://schema.org", Type: "CreativeWork"}}
}
