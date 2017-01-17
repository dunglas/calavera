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

	"github.com/dunglas/calavera/extractor"
	"github.com/dunglas/calavera/schema"
)

const FILE_PERMS = 0644
const DIR_PERMS = 0755

func main() {
	inputPath, outputPath, prettify := parseFlags()

	var files []string
	var extractors = []extractor.Extractor{extractor.NewMarkdown(inputPath)}

	if ge, err := extractor.NewGit(inputPath); nil == err {
		extractors = append(extractors, ge)
	} else {
		log.Println(`"` + inputPath + `" is not a Git repository. Authors and date metadata will NOT be extracted.`)
	}

	walkFunc := func(path string, _ os.FileInfo, err error) error {
		if nil != err || !strings.HasSuffix(path, ".md") {
			return nil
		}

		abs, err := filepath.Abs(path)
		check(err)
		rel, err := filepath.Rel(inputPath, abs)
		check(err)
		files = append(files, rel)

		return nil
	}

	if err := filepath.Walk(inputPath, walkFunc); nil != err {
		check(err)
	}

	entrypoint := schema.NewItemList()
	for _, file := range files {
		// Cannot use a go routine because src-d/go-git isn't thread safe
		convert(file, outputPath, extractors, prettify)
		entrypoint.Element = append(entrypoint.Element, getOutputPath(file))
	}

	check(ioutil.WriteFile(outputPath+"/_index.jsonld", marshal(entrypoint, prettify), FILE_PERMS))
}

func marshal(v interface{}, prettify bool) []byte {
	var jsonContent []byte
	var err error
	if prettify {
		jsonContent, err = json.MarshalIndent(v, "", "\t")
	} else {
		jsonContent, err = json.Marshal(v)
	}
	check(err)

	return jsonContent
}

func check(err error) {
	if nil == err {
		return
	}

	log.Fatalln(err)
	panic(err)
}

func parseFlags() (string, string, bool) {
	flag.Usage = func() {
		fmt.Println("calavera input_directory output_directory")
	}

	prettify := flag.Bool("prettify", false, "Prettify json output")

	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalln("Input and output directories are mandatory arguments.")
	}

	inputPath, err := filepath.Abs(flag.Arg(0))
	check(err)

	outputPath, err := filepath.Abs(flag.Arg(1))
	check(err)

	return inputPath, outputPath, *prettify
}

func convert(path string, outputDirectory string, extractors []extractor.Extractor, prettify bool) {
	creativeWork := schema.NewCreativeWork()

	for _, extractor := range extractors {
		err := extractor.Extract(creativeWork, path)
		check(err)
	}

	jsonContent := marshal(creativeWork, prettify)

	outputPath := outputDirectory + "/" + getOutputPath(path)
	outputSubdirectory := filepath.Dir(outputPath)

	err := os.MkdirAll(outputSubdirectory, DIR_PERMS)
	check(err)

	err = ioutil.WriteFile(outputPath, jsonContent, FILE_PERMS)
	check(err)
}

func getOutputPath(originalPath string) string {
	return originalPath[:len(originalPath)-3] + ".jsonld"
}
