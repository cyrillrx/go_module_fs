package main

import (
	"os"
	"time"
)

type File struct {
	Name    string    `json:"name,omitempty"`
	Size    int64     `json:"size,omitempty"`
	ModTime time.Time `json:"time,omitempty"`
	IsDir   bool      `json:"is_dir,omitempty"`
}

func makeFile(src os.FileInfo) File {

	var f File

	f.Name = src.Name()
	f.Size = src.Size()
	f.ModTime = src.ModTime()
	f.IsDir = src.IsDir()

	return f
}
