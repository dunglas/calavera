package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dunglas/calavera/extractor"
	"github.com/dunglas/calavera/schema"
)

func main() {	
	flag.Usage = func() {
		fmt.Println("calavera input_directory output_directory")
	}

	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalln("Input and output directories are mandatory arguments.")
	}
	
	var files []string
	var extractors = []extractor.Extractor{extractor.Markdown{}, extractor.Git{}}
	
	walkFunc := func(path string, _ os.FileInfo, err error) error {
		if nil == err && strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}

		return nil
	}

	if err := filepath.Walk(flag.Arg(0), walkFunc); nil != err {
		check(err)
	}

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			convert(file, flag.Arg(1), extractors)
			defer wg.Done()
		}(file)
	}

	wg.Wait()
}

func convert(path string, outputDirectory string, extractors []extractor.Extractor) {
	creativeWork := schema.NewCreativeWork()

	for _, extractor := range extractors {
		err := extractor.Extract(creativeWork, path)
		check(err)
	}

	jsonContent, err := json.MarshalIndent(creativeWork, "", "\t")
	check(err)

	outputPath := fmt.Sprint(outputDirectory, "/", path[:len(path) - 3], ".jsonld")
	outputSubdirectory := filepath.Dir(outputPath)

	err = os.MkdirAll(outputSubdirectory, 0755)
	check(err)

	err = ioutil.WriteFile(outputPath, jsonContent, 0644)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
