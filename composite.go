package main

import "fmt"

type Entry interface {
	Name() string
	Size() int64
	PrintList(prefix string)
}

type File struct {
	name string
	size int64
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int64 {
	return f.size
}

func (f *File) PrintList(prefix string) {
	fmt.Println(prefix + "/" + f.Name() + " (" + fmt.Sprint(f.Size()) + ")")
}

func NewFile(name string, size int64) *File {
	return &File{name: name, size: size}
}

type Directory struct {
	name    string
	entries []Entry
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name, entries: []Entry{}}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int64 {
	var total int64 = 0
	for _, entry := range d.entries {
		total += entry.Size()
	}
	return total
}

func (d *Directory) Add(entry Entry) {
	d.entries = append(d.entries, entry)
}

func (d *Directory) PrintList(prefix string) {
	for _, entry := range d.entries {
		entry.PrintList(prefix + "/" + d.Name())
	}
}

func ExecComposite() {
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	usrDir := NewDirectory("usr")
	homeDir := NewDirectory("home")
	rootDir.Add(binDir)
	rootDir.Add(usrDir)
	rootDir.Add(homeDir)

	binDir.Add(NewFile("vi", 10000))
	binDir.Add(NewFile("latex", 20000))

	usrDir.Add(NewFile("memo.txt", 3000))
	usrDir.Add(NewFile("game.exe", 40000))

	homeDir.Add(NewFile("diary.html", 5000))
	homeDir.Add(NewFile("photo.jpg", 6000))

	rootDir.PrintList("")
}
