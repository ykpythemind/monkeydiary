package main

import (
	"io"
	"io/ioutil"
	"os"
)

func createFile(src io.Reader) (file *os.File, err error) {
	tmpfile, err := ioutil.TempFile("", "diary")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(tmpfile, src)
	if err != nil {
		tmpfile.Close()
		return nil, err
	}

	return tmpfile, nil
}
