package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGithubExecute(t *testing.T) {
	github := newGithubService()

	data := "aaaaaaaa"

	file, err := createFile(strings.NewReader(data))
	if err != nil {
		t.Errorf("create file fail %s", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	fmt.Println(file.Name())

	if err := github.execute(file.Name()); err != nil {
		t.Error(err)
	}
}
