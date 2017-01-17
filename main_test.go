package main

import (
	"encoding/json"
	"github.com/dunglas/calavera/schema"
	"io/ioutil"
	"os"
	"testing"
)

func exists(path string, t *testing.T) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error(`The file "` + path + `" must exist.`)
	}
}

func TestFoo(t *testing.T) {
	os.RemoveAll("out/")
	os.Args = append(os.Args, "fixtures", "out")

	main()

	exists("out/_index.jsonld", t)
	exists("out/index.jsonld", t)
	exists("out/foo.jsonld", t)
	exists("out/level1/bar.jsonld", t)

	bar, _ := ioutil.ReadFile("out/level1/bar.jsonld")
	barDoc := &schema.CreativeWork{}
	json.Unmarshal(bar, barDoc)

	if "level1/bar.jsonld" != barDoc.Id {
		t.Error(`The "@id" property of "level1/bar.jsonld" must be set.`)
	}

	index, _ := ioutil.ReadFile("out/_index.jsonld")
	indexDoc := &schema.CreativeWork{}
	json.Unmarshal(index, indexDoc)

	if "_index.jsonld" != indexDoc.Id {
		t.Error(`The "@id" property of "_index.jsonld" must be set.`)
	}
}
