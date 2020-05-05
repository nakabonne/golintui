package golangcilint

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	tmpGoFileName     = "main.go"
	tmpConfigFileName = ".golangci.yml"
)

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

// tmpConfigFile creates a temporary config file and returns its path.
func tmpConfigFile(yml []byte) (string, func(), error) {
	tmpDir, err := ioutil.TempDir("", "nakabonne-golintui")
	if err != nil {
		return "", nil, err
	}
	path := filepath.Join(tmpDir, tmpConfigFileName)
	_, err = create(path, yml)
	if err != nil {
		return "", nil, err
	}

	cleaner := func() {
		os.RemoveAll(tmpDir)
	}
	return path, cleaner, nil

}

func create(path string, contents []byte) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if _, err = f.Write([]byte(contents)); err != nil {
		return nil, err
	}
	return f, nil
}
