package schema

// JsonLd represents a JSON-LD document wiht "@context" and "@type" keys.
type JsonLd struct {
	Context string `json:"@context,omitempty"`
	Type string `json:"@type,omitempty"`
}
