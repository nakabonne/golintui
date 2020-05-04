package golangcilint

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const mainContents = `package main

func main() {
}`

// tmpProject creates a temporary Go project.
func tmpProject() (string, func(), error) {
	tmpDir, err := ioutil.TempDir("", "nakabonne-golintui")
	if err != nil {
		return "", nil, err
	}
	_, err = create(filepath.Join(tmpDir, "main.go"), mainContents)
	if err != nil {
		return "", nil, err
	}
	cleaner := func() {
		os.RemoveAll(tmpDir)
	}
	return tmpDir, cleaner, nil
}

func create(path, contents string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if _, err = f.Write([]byte(contents)); err != nil {
		return nil, err
	}
	return f, nil
}
