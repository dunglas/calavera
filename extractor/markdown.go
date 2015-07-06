package extractor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/dunglas/calavera/schema"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

type Markdown struct {
}

func (markdown Markdown) Extract(creativeWork *schema.CreativeWork, path string) error {
	markdownContent, err := ioutil.ReadFile(path)
	if nil != err {
		return err
	}

	html := blackfriday.MarkdownCommon(markdownContent)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if nil != err {
		return err
	}

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		url, _ := url.Parse(link)

		if !url.IsAbs() && strings.HasSuffix(link, ".md") {
			s.SetAttr("href", fmt.Sprint(link[:len(link) - 3], ".jsonld"))
		}
	})

	creativeWork.Name = doc.Find("h1").Text()
	creativeWork.Text, err = doc.Find("body").Html()
	if nil != err {
		return err
	}

	return nil
}
