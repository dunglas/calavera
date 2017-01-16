package extractor

import (
	"bufio"
	"log"
	"os/exec"
	"strings"

	"github.com/dunglas/calavera/schema"
)

var gitPath string

func init() {
	var err error

	gitPath, err = exec.LookPath("git")
	if nil != err {
		log.Fatalln("git is not available in the PATH. Install it to extract git metadata.")
	}
}

// Git extracts metadata from the Git repository containing Markdown files.
type Git struct {
}

// Extract extracts the list of contributors to the file, and date of modifications.
func (git Git) Extract(creativeWork *schema.CreativeWork, path string) error {
	if "" == gitPath {
		return nil
	}

	cmd := exec.Command(gitPath, "log", "--format=%an;%ae;%aI", path)
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		parts := strings.Split(strings.TrimSpace(scanner.Text()), ";")

		author := schema.NewPerson(parts[0], parts[1])
		creativeWork.Author = append([]schema.Person{*author}, creativeWork.Author...)

		creativeWork.DateCreated = parts[2]
		if "" == creativeWork.DateModified {
			creativeWork.DateModified = parts[2]
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalln("You are not in a git repository.")
		return err
	}

	return nil
}
