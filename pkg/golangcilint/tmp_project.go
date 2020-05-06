package golangcilint

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const tmpGoFileName = "main.go"

var mainContents = []byte(`package main

func main() {
}`)

// tmpProject creates a temporary Go project.
func tmpProject() (string, func(), error) {
	tmpDir, err := ioutil.TempDir(".", "nakabonne-golintui")
	if err != nil {
		return "", nil, err
	}
	_, err = create(filepath.Join(tmpDir, tmpGoFileName), mainContents)
	if err != nil {
		return "", nil, err
	}

	cleaner := func() {
		os.RemoveAll(tmpDir)
	}
	return tmpDir, cleaner, nil
}

func create(path string, contents []byte) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if _, err = f.Write(contents); err != nil {
		return nil, err
	}
	return f, nil
}
