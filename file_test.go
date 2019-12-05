package main

import (
	"os"
	"strings"
	"testing"
)

func TestCreateFile(t *testing.T) {
	data := "aaaaaaaa"

	file, err := createFile(strings.NewReader(data))
	if err != nil {
		t.Errorf("create file fail %s", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()
}
