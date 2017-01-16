package extractor

import (
	"github.com/dunglas/calavera/schema"
)

// Extractor updates the schema.CreativeWork structure.
type Extractor interface {
	Extract(creativeWork *schema.CreativeWork, path string) error
}
