package extractor

import (
	"testing"
	"github.com/dunglas/calavera/schema"
	"strings"
)

func TestMarkdown_Extract(t *testing.T) {
	creativeWork := schema.NewCreativeWork()

	extractor := Markdown{}
	err := extractor.Extract(creativeWork, "../fixtures/foo.md")

	if err != nil {
		t.Error(err)
	}

	if creativeWork.Name != "Foo" {
		t.Errorf("Title should be \"Foo\", but is \"%s\"." + creativeWork.Name)
	}

	if strings.Contains(creativeWork.Text, ".md") {
		t.Error("References to Markdown file must be changed to references to JSON-LD files.")
	}

	if strings.Contains(creativeWork.Text, "rel=\"nofollow\"") {
		t.Error("Links must be followed by spiders.")
	}

	if !strings.Contains(creativeWork.Text, "class=\"language-php\"") {
		t.Error("Classes must be preserved.")
	}
}

func TestGit_Extract(t *testing.T) {
	creativeWork := schema.NewCreativeWork()

	extractor := Git{}
	err := extractor.Extract(creativeWork, "../fixtures/foo.md")

	if err != nil {
		t.Error(err)
	}

	if "" == creativeWork.DateModified {
		t.Error("The creation date must be extracted.")
	}

	if "" == creativeWork.DateModified {
		t.Error("The modifiation date must be extracted.")
	}

	var found = false
	for _, person := range creativeWork.Author {
		if person.Email == "dunglas@gmail.com" && person.Name == "Kévin Dunglas" {
			found = true
		}
	}

	if !found {
		t.Error("Kévin must be part of the authors.")
	}
}
