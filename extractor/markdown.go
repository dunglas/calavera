package extractor

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dunglas/calavera/schema"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"regexp"
)

// Markdown is an extractor converting the content of .md files to HTML.
type Markdown struct {
}

// Extract converts the Markdown syntax to HTML in a secure way.
// The generated file is sanitized: all special tags and characters stripped (JavaScript, CSS...) or escaped.
// A "class" attributes containing language indication for syntax highlighting is added to all code snippets.
// All references to ".md" files are converted to links pointing to ".jsonld" files.
func (markdown Markdown) Extract(creativeWork *schema.CreativeWork, path string) error {
	markdownContent, err := ioutil.ReadFile(path)
	if nil != err {
		return err
	}

	unsafe := blackfriday.MarkdownCommon(markdownContent)
	p := bluemonday.UGCPolicy()
	p.RequireNoFollowOnLinks(false)
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	html := p.SanitizeBytes(unsafe)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if nil != err {
		return err
	}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		url, _ := url.Parse(link)

		if !url.IsAbs() {
			s.SetAttr("href", strings.Replace(link, ".md", ".jsonld", 1))
		}
	})

	creativeWork.Name = doc.Find("h1").Text()
	creativeWork.Text, err = doc.Find("body").Html()
	if nil != err {
		return err
	}

	return nil
}
