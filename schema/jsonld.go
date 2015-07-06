package schema

type JsonLd struct {
	Context string `json:"@context,omitempty"`
	Type string `json:"@type,omitempty"`
}
