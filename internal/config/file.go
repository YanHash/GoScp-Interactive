package config

import (
	"log"
	"os"
	"path/filepath"
)

// can be folder or file
type Element struct {
	path        string
	name        string
	isFolder    bool
	subElements []Element
}

func DescribeFile(path string, name string) *Element {
	return &Element{
		path:        path,
		name:        name,
		isFolder:    false,
		subElements: nil,
	}
}

func DescribeFolder(path string, name string) *Element {
	return &Element{
		path:        path,
		name:        name,
		isFolder:    false,
		subElements: nil,
	}
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		log.Fatal("an error occured while checking if the path provided is a Directory - ", err)
	}
	return info.IsDir()
}

func ListElements(path string) []Element {
	isDir := isDirectory(path)
	output = []Element
	if isDir {
		elements, err := os.ReadDir(path)
		if err != nil {
			log.Fatal("an error occured while listing elements - ", err)
		}

		for _, element := range elements {
			e = Element{
				path: _,
				name: e.name,
			}
		}

	}
}

type File struct {
	path     string
	fileName string
}

func (f *File) GetPath() string {
	return f.path
}

func (f *File) GetFilePath() string {
	return filepath.Join(f.path, f.fileName)
}

func NewFile(path, fileName string) *File {
	if path == "" || fileName == "" {
		panic("path and fileName cannot be empty")
	}
	return &File{path: path, fileName: fileName}
}
