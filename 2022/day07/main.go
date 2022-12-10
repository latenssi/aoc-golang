package main

import (
	"strconv"
	"strings"
)

const defaultTotalSize = 70000000
const spaceRequiredForUpdate = 30000000

type FileSystem struct {
	root       *Directory
	currentDir *Directory
	totalSize  int
}

func NewFileSystem() *FileSystem {
	root := &Directory{name: "/"}
	return &FileSystem{
		root:       root,
		currentDir: root,
		totalSize:  defaultTotalSize,
	}
}

func (fs *FileSystem) ChangeDir(path string) {
	switch path {
	case "/":
		fs.currentDir = fs.root
	case "..":
		fs.currentDir = fs.currentDir.parent
	default:
		for _, dir := range fs.currentDir.subdirs {
			if dir.name == path {
				fs.currentDir = dir
				return
			}
		}
		panic("Directory not found: " + path)
	}
}

func (fs *FileSystem) MakeDir(path string) {
	dir := &Directory{name: path, parent: fs.currentDir}
	fs.currentDir.subdirs = append(fs.currentDir.subdirs, dir)
}

func (fs *FileSystem) AddFile(path string, size int) {
	file := &File{name: path, size: size}
	fs.currentDir.files = append(fs.currentDir.files, file)
}

func (fs *FileSystem) AllDirs() []*Directory {
	dirs := []*Directory{fs.root}
	dirs = append(dirs, fs.root.SubDirs(true)...)
	return dirs
}

func (fs *FileSystem) TotalSize() int {
	return fs.totalSize
}

func (fs *FileSystem) FreeSpace() int {
	return fs.totalSize - fs.root.Size()
}

type Directory struct {
	name    string
	parent  *Directory
	subdirs []*Directory
	files   []*File
}

func (d *Directory) Size() int {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, subdir := range d.subdirs {
		size += subdir.Size()
	}
	return size
}

func (d *Directory) SubDirs(recursive bool) []*Directory {
	dirs := []*Directory{}
	for _, subdir := range d.subdirs {
		dirs = append(dirs, subdir)
		if recursive {
			dirs = append(dirs, subdir.SubDirs(recursive)...)
		}
	}
	return dirs
}

type File struct {
	name string
	size int
}

func Day(input string, part int) int {
	fs := NewFileSystem()

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if line[0] == '$' {
			// Handle commands
			if parts[1] == "cd" {
				fs.ChangeDir(parts[2])
			}
		} else {
			// Handle ls output
			if parts[0] == "dir" {
				// Add directory
				fs.MakeDir(parts[1])
			} else {
				// Add file
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}
				fs.AddFile(parts[1], size)
			}
		}
	}

	if part == 1 {
		sum := 0
		for _, dir := range fs.AllDirs() {
			size := dir.Size()
			if size <= 100000 {
				sum += size
			}
		}
		return sum
	}

	if part == 2 {
		spaceNeeded := spaceRequiredForUpdate - fs.FreeSpace()
		min := fs.TotalSize()
		for _, dir := range fs.AllDirs() {
			size := dir.Size()
			if size >= spaceNeeded && size < min {
				min = size
			}
		}
		return min
	}

	return 0
}
