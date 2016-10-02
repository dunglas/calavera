package extractor

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/dunglas/calavera/schema"
	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"regexp"
)

type Markdown struct {
}

func (markdown Markdown) Extract(creativeWork *schema.CreativeWork, path string) error {
	markdownContent, err := ioutil.ReadFile(path)
	if nil != err {
		return err
	}

	unsafe := blackfriday.MarkdownCommon(markdownContent)
	p := bluemonday.UGCPolicy()
	p.RequireNoFollowOnLinks(false)
	p.AllowAttrs("class").Matching(regexp.MustCompile("^[a-zA-Z0-9-_]+$")).Globally()
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
