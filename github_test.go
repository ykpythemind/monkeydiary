package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type githubServiceImplMock struct {
}

func (g *githubServiceImplMock) clone() (dir string, err error) {
	dir, err = ioutil.TempDir("", "test")
	if err != nil {
		panic(err)
	}

	// dataディレクトリも必要
	if err := os.Mkdir(filepath.Join(dir, "data"), 0777); err != nil {
		panic(err)
	}

	return
}

func (g *githubServiceImplMock) add(path string) error {
	return nil
}

func (g *githubServiceImplMock) commit() error {
	return nil
}

func (g *githubServiceImplMock) push() error {
	return nil
}

func TestExecuteGitOperation(t *testing.T) {
	github := &githubServiceImplMock{}

	data := "aaaaaaaa"

	file, err := createFile(strings.NewReader(data))
	if err != nil {
		t.Errorf("create file fail %s", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	if err := executeGitOperation(file.Name(), github); err != nil {
		t.Error(err)
	}
}
