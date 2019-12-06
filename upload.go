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

type uploadService interface {
	upload(source io.Reader) error
}

type uploadGithubService struct {
	config *config
}

func newUploadGithubService(config *config) *uploadGithubService {
	return &uploadGithubService{config: config}
}

// upload executes git clone, copy specified file, commit, and push.
func executeUploadService(source io.Reader, uploadService uploadService) error {
	log.Println("uploading")
	if err := uploadService.upload(source); err != nil {
		return err
	}
	log.Println("done")
	return nil
}

func (g *uploadGithubService) fetchGithub() (repo *git.Repository, dir string, err error) {
	// Tempdir to clone the repository
	dir, err = ioutil.TempDir("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.RemoveAll(dir) // clean up

	// Clones the repository into the given dir, just as a normal git clone does
	repo, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:   g.config.diaryRepository,
		Auth:  g.newAuth(),
		Depth: 1,
	})
	if err != nil {
		return
	}

	return
}

func (g *uploadGithubService) upload(source io.Reader) error {
	log.Println("fetch current data")
	repo, dir, err := g.fetchGithub()
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	newFilePath := filepath.Join(dir, "data", time.Now().Format(time.RFC3339))

	// Create new file
	newFile, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	log.Println("copy src file to repo's directory")
	_, err = io.Copy(newFile, source)
	if err != nil {
		log.Fatal(err)
	}
	gitAddTarget := filepath.Join("data", filepath.Base(newFilePath))

	if _, err := worktree.Add(gitAddTarget); err != nil {
		return err
	}
	log.Println("add")

	_, err = worktree.Commit("Monkey wrote diary.", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "ykpythemind",
			Email: "yukibukiyou@gmail.com",
			When:  time.Now(),
		},
	})
	log.Println("commit")

	if err := repo.Push(&git.PushOptions{
		Auth: g.newAuth(),
	}); err != nil {
		return err
	}
	log.Println("push")

	return nil
}

func (g *uploadGithubService) newAuth() *http.BasicAuth {
	return &http.BasicAuth{
		Username: g.config.userName,
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
