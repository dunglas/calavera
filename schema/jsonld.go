package schema

// JsonLd represents a JSON-LD document with "@context" and "@type" keys.
type JsonLd struct {
	Context string `json:"@context,omitempty"`
	Type    string `json:"@type,omitempty"`
}
