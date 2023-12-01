package day7

import (
	"fmt"
)

type file struct {
	name string
	size int
}

type dir struct {
	dirs  map[string]*dir
	files map[string]int
	size  int
}

type FileSystem struct {
	root    *dir
	currDir *dir
	wd      *Stack
}

func NewFS() *FileSystem {
	rootWd := &Stack{path: []string{"/"}, length: 1}
	rootDir := &dir{dirs: make(map[string]*dir), files: make(map[string]int)}
	// rootDir := &dir{dirs: make(map[string]*dir), files: []file{}}

	return &FileSystem{
		root:    rootDir,
		currDir: rootDir,
		wd:      rootWd,
	}
}

func (fs *FileSystem) CDInto(d string) {
	if d == "/" {
		fs.currDir = fs.root
		fs.wd.reset()
		return
	}

	if _, exists := fs.currDir.dirs[d]; !exists {
		fs.currDir.dirs[d] = &dir{dirs: make(map[string]*dir), files: make(map[string]int)}
		// fs.currDir.dirs[d] = &dir{dirs: make(map[string]*dir), files: []file{}}
	}

	fs.currDir = fs.currDir.dirs[d]
	fs.wd.insert(d)
}

func (fs *FileSystem) CDBack() {
	if fs.wd.length == 1 {
		return
	}

	err := fs.wd.pop()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fs.currDir = fs.root

	for i, dir := range fs.wd.path {
		if i == 0 {
			continue
		}
		fs.currDir = fs.currDir.dirs[dir]
	}
}

func (fs *FileSystem) GetFile(name string, size int) {
	if _, exists := fs.currDir.files[name]; !exists {
		fs.currDir.files[name] = size
		fs.currDir.size += size
	}
}
