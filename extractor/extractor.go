package extractor

import (
	"github.com/dunglas/calavera/schema"
)

type Extractor interface {
	Extract(creativeWork *schema.CreativeWork, path string) error
}
