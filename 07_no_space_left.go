package adventofcode2022

import (
	"math"
	"strconv"
	"strings"
)

type Directory struct {
	Name            string
	Size            int
	ParentDirectory *Directory
	SubDirectories  []*Directory
	Files           []*File
}

type File struct {
	Name string
	Size int
}

func (d *Directory) addDirectory(name string) *Directory {
	subd := d.getSubdirectory(name)
	if subd == nil {
		new := NewDirectory(name, d)
		d.SubDirectories = append(d.SubDirectories, new)
		return new
	}
	return subd
}

func (d *Directory) addFile(file *File) {
	d.Files = append(d.Files, file)
	d.Size += file.Size

	for d.ParentDirectory != nil {
		d = d.ParentDirectory
		d.Size += file.Size
	}
}

func (d *Directory) getSubdirectory(name string) *Directory {
	for _, subDirectory := range d.SubDirectories {
		if subDirectory.Name == name {
			return subDirectory
		}
	}
	return nil
}

func NewDirectory(name string, parentDirectory *Directory) *Directory {
	return &Directory{
		Name:            name,
		ParentDirectory: parentDirectory,
		SubDirectories:  []*Directory{},
		Files:           []*File{},
	}
}

func NewFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

func executeInstructions(root *Directory, instructions []string) *Directory {
	cwd := root

	for _, cmd := range instructions {
		switch {
		case cmd == "$ cd /", cmd == "$ ls", cmd == "dir":
			continue // Commands not used
		case strings.HasPrefix(cmd, "$ cd"):
			dirName := strings.Split(cmd, " ")[2]
			switch dirName {
			case "..":
				cwd = cwd.ParentDirectory
			default:
				cwd = cwd.addDirectory(dirName)
			}
		default:
			vals := strings.Split(cmd, " ")
			size, _ := strconv.Atoi(vals[0])
			cwd.addFile(NewFile(vals[1], size))
		}
	}

	return root
}

func sizeOfDirectoriesUnderThreshold(d *Directory, threshold int) int {
	total := 0
	if d.Size <= threshold {
		total += d.Size
	}

	for _, subDirectory := range d.SubDirectories {
		total += sizeOfDirectoriesUnderThreshold(subDirectory, threshold)
	}

	return total
}

func sizeOfDirectoryClosestOverSize(d *Directory, requiredSize int, currentCandidate int) int {
	if d.Size >= requiredSize && d.Size <= currentCandidate {
		currentCandidate = d.Size
	}

	for _, subDirectory := range d.SubDirectories {
		res := sizeOfDirectoryClosestOverSize(subDirectory, requiredSize, currentCandidate)
		if res >= requiredSize && res <= currentCandidate {
			currentCandidate = res
		}
	}

	return currentCandidate
}

func daySevenChallengeOne(readout []string) int {
	root := executeInstructions(NewDirectory("/", nil), readout)
	total := sizeOfDirectoriesUnderThreshold(root, 100000)
	return total
}

func daySevenChallengeTwo(readout []string) int {
	root := executeInstructions(NewDirectory("/", nil), readout)
	target := 30000000 - (70000000 - root.Size)
	return sizeOfDirectoryClosestOverSize(root, target, math.MaxInt)
}
