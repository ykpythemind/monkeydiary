package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type githubService interface {
	clone() (dir string, err error)
	add(path string) error
	commit() error
	push() error
}

type githubServiceImpl struct {
	repository *git.Repository
}

func newGithubService() githubService {
	return &githubServiceImpl{}
}

// executeGitOperation executes git clone, copy specified file, commit, and push.
func executeGitOperation(srcFilePath string, githubService githubService) error {
	dir, err := githubService.clone()
	if err != nil {
		return err
	}

	log.Println("copy src file to repo's directory")
	sourceFile, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	newFilePath := filepath.Join(dir, "data", time.Now().Format(time.RFC3339))

	// Create new file
	newFile, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("add")
	if err := githubService.add(filepath.Join("data", filepath.Base(newFilePath))); err != nil {
		return err
	}

	log.Println("commit")
	if err := githubService.commit(); err != nil {
		return err
	}

	log.Println("push")
	if err := githubService.push(); err != nil {
		return err
	}

	log.Println("done")
	return nil
}

func (g *githubServiceImpl) clone() (dir string, err error) {
	// Tempdir to clone the repository
	dir, err = ioutil.TempDir("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.RemoveAll(dir) // clean up

	// Clones the repository into the given dir, just as a normal git clone does
	repository, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:   "https://github.com/ykpythemind/monkeydiary_web",
		Auth:  newAuth(),
		Depth: 1,
	})

	if err != nil {
		return
	}

	g.repository = repository

	return
}

func (g *githubServiceImpl) add(path string) error {
	worktree, err := g.repository.Worktree()
	if err != nil {
		return err
	}

	if _, err := worktree.Add(path); err != nil {
		return err
	}

	return nil
}

func (g *githubServiceImpl) commit() error {
	worktree, err := g.repository.Worktree()
	if err != nil {
		return err
	}

	_, err = worktree.Commit("Monkey wrote diary.", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "ykpythemind",
			Email: "yukibukiyou@gmail.com",
			When:  time.Now(),
		},
	})

	return err
}

func (g *githubServiceImpl) push() error {
	err := g.repository.Push(&git.PushOptions{
		Auth: newAuth(),
	})

	if err != nil {
		return err
	}

	return nil
}

func newAuth() *http.BasicAuth {
	return &http.BasicAuth{
		Username: "ykpythemind",
		Password: accessToken(),
	}

}

func accessToken() string {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		log.Fatal("set GITHUB_TOKEN env")
	}

	return token
}
