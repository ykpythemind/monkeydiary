package main

import (
	"io"
	"strings"
	"testing"
)

type uploadServiceMock struct {
}

func (g *uploadServiceMock) upload(source io.Reader) error {
	return nil
}

func TestExecuteUploadService(t *testing.T) {
	service := &uploadServiceMock{}

	data := "aaaa"

	if err := executeUploadService(strings.NewReader(data), service); err != nil {
		t.Error(err)
	}
}

func TestUploadGithubService(t *testing.T) {
	config := &config{
		diaryRepository: "https://github.com/ykpythemind/test_of_monkeydiary",
		userName:        "ykpythemind",
	}
	service := newUploadGithubService(config)

	src := strings.NewReader("aioaio")

	if err := service.upload(src); err != nil {
		t.Error(err)
	}
}
