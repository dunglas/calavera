package main

import (
	"os"
	"testing"
)

func TestFoo(t *testing.T) {
	os.RemoveAll("out/")
	os.Args = append(os.Args, "fixtures", "out")

	main()

	if _, err := os.Stat("out/index.jsonld"); os.IsNotExist(err) {
		t.Error("The file \"out/index.jsonld\" must exist.")
	}

	if _, err := os.Stat("out/foo.jsonld"); os.IsNotExist(err) {
		t.Error("The file \"out/foo.jsonld\" must exist.")
	}

	if _, err := os.Stat("out/level1/bar.jsonld"); os.IsNotExist(err) {
		t.Error("The file \"out/level1/bar.jsonld\" must exist.")
	}
}
