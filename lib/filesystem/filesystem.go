package filesystem

type Shell struct {
	fs         *FileSystem
	currentDir *Directory
}

func NewShell(fs *FileSystem) *Shell {
	return &Shell{
		fs:         fs,
		currentDir: fs.root,
	}
}

func (s *Shell) ChangeDir(path string) {
	switch path {
	case "/":
		s.currentDir = s.fs.root
	case "..":
		s.currentDir = s.currentDir.parent
	default:
		for _, dir := range s.currentDir.subdirs {
			if dir.name == path {
				s.currentDir = dir
				return
			}
		}
		panic("Directory not found: " + path)
	}
}

func (s *Shell) MakeDir(name string) {
	s.fs.makeDir(name, s.currentDir)
}

func (s *Shell) AddFile(name string, size int) {
	s.fs.addFile(name, size, s.currentDir)
}

type FileSystem struct {
	root      *Directory
	totalSize int
}

func NewFileSystem(size int) *FileSystem {
	return &FileSystem{
		root:      &Directory{name: "/"},
		totalSize: size,
	}
}

// makeDir is a kernel method
func (fs *FileSystem) makeDir(name string, parent *Directory) {
	dir := &Directory{name: name, parent: parent}
	parent.subdirs = append(parent.subdirs, dir)
}

// addFile is a kernel method
func (fs *FileSystem) addFile(name string, size int, parent *Directory) {
	file := &File{name: name, size: size}
	parent.files = append(parent.files, file)
}

// AllDirs is a public method
func (fs *FileSystem) AllDirs() []*Directory {
	dirs := []*Directory{fs.root}
	dirs = append(dirs, fs.root.SubDirs(true)...)
	return dirs
}

// TotalSize is a public method
func (fs *FileSystem) TotalSize() int {
	return fs.totalSize
}

// FreeSpace is a public method
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
