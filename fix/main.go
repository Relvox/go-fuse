package main

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	filepath.Walk("..", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.Contains(path, "_linux") {
			return nil
		}
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		newPath := strings.ReplaceAll(path, "_linux", "_windows")
		buf = []byte(strings.ReplaceAll(string(buf), "build linux", "build windows"))
		return ioutil.WriteFile(newPath, buf, 0644)
	})
}
