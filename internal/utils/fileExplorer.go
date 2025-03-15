package utils

import (
	"log"
	"os"
)

type FileExplorer struct {
	workingPath string
}

func (f *FileExplorer) GoToCurrentPath() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f.workingPath = dir
}

func (f *FileExplorer) GoToPath(path string) {
	if path == "" {
		return
	}
	f.workingPath = path
}
