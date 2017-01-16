package main

import (
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
}
