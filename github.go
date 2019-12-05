package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type githubService interface {
	execute(srcFilePath string) error
}

type githubServiceImpl struct {
}

func newGithubService() githubService {
	return githubServiceImpl{}
}

// Execute executes git clone, copy specified file, commit, and push.
func (g githubServiceImpl) execute(srcFilePath string) error {
	repo, dir, err := clone()
	if err != nil {
		return err
	}

	// prev, err := filepath.Abs(".")
	// if err != nil {
	// 	return err
	// }
	// defer os.Chdir(prev)

	// os.Chdir(dir)

	sourceFile, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create(filepath.Join(dir, time.Now().Format(time.RFC822)))
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	fmt.Println("worktree add")

	_, err = worktree.Add(newFile.Name())
	if err != nil {
		return err
	}

	fmt.Println("commit")
	_, err = worktree.Commit("ya", &git.CommitOptions{})
	if err != nil {
		return err
	}

	fmt.Println("push")
	if err := repo.Push(&git.PushOptions{}); err != nil {
		return err
	}

	return nil
}

func clone() (repository *git.Repository, dir string, err error) {
	// Tempdir to clone the repository
	dir, err = ioutil.TempDir("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.RemoveAll(dir) // clean up

	// Clones the repository into the given dir, just as a normal git clone does
	repository, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: "https://github.com/ykpythemind/monkeydiary_web",
		Auth: &http.BasicAuth{
			Username: "ykpythemind",
			Password: accessToken(),
		},
		Depth: 1,
	})

	if err != nil {
		return
	}

	return
}

func accessToken() string {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		log.Fatal("set GITHUB_TOKEN env")
	}

	return token
}
