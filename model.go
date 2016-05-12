package main

import (
	"encoding/json"
	"fmt"
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

func (f File) Print() {

	b, err := json.Marshal(f)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func (f File) PrettyPrint() {

	b, err := json.MarshalIndent(f, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
