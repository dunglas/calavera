package extractor

import (
	"errors"
	"github.com/dunglas/calavera/schema"
	gogit "gopkg.in/src-d/go-git.v4"
	"os"
	"path/filepath"
	"strings"
)

// Git extracts metadata from the Git repository containing Markdown files.
type Git struct {
	inputDirectory string
	gitDirectory   string
	repository     *gogit.Repository
}

func findGitDir(path string) (string, error) {
	if _, err := os.Stat(path + "/.git/config"); err == nil {
		return path + "/.git", nil
	}

	parentDir := filepath.Dir(path)
	if strings.HasSuffix(parentDir, "/") {
		return "", errors.New("No Git repository found")
	}

	return findGitDir(parentDir)
}

// NewGit returns a new instance properly configured of the Git extractor
func NewGit(inputDirectory string) (*Git, error) {
	var err error
	var gitDirectory string

	if gitDirectory, err = findGitDir(inputDirectory); nil != err {
		return nil, err
	}

	if r, err := gogit.NewFilesystemRepository(gitDirectory); nil == err {
		return &Git{
			inputDirectory: inputDirectory,
			gitDirectory:   gitDirectory,
			repository:     r,
		}, nil
	}

	return nil, err
}

// Extract extracts the list of contributors to the file, and date of modifications.
func (git Git) Extract(creativeWork *schema.CreativeWork, path string) error {
	path, _ = filepath.Rel(filepath.Dir(git.gitDirectory), git.inputDirectory+"/"+path)

	ref, err := git.repository.Head()
	if nil != err {
		return err
	}

	c, err := git.repository.Commit(ref.Hash())
	if nil != err {
		return err
	}

	revs, err := gogit.References(c, path)
	if err != nil {
		return err
	}

	for _, v := range revs {
		if !authorExists(creativeWork.Author, v.Author.Email) {
			author := schema.NewPerson(v.Author.Name, v.Author.Email)
			creativeWork.Author = append([]schema.Person{*author}, creativeWork.Author...)
		}

		creativeWork.DateModified = &v.Author.When
		if nil == creativeWork.DateCreated {
			creativeWork.DateCreated = &v.Author.When
		}
	}

	return nil
}

// authorExists tests if an author is already in the list
func authorExists(authors []schema.Person, email string) bool {
	for _, p := range authors {
		if email == p.Email {
			return true
		}
	}

	return false
}
