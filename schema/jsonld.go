package schema

// JsonLd represents a JSON-LD document with "@context" and "@type" keys.
type JsonLd struct {
	Id      string `json:"@id,omitempty"`
	Context string `json:"@context,omitempty"`
	Type    string `json:"@type,omitempty"`
}
