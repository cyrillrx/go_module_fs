package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

func main() {

	args := os.Args[1:]
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "./"
	}

	var output []File

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		output = append(output, makeFile(f))
		fmt.Println(f.Name())
	}

	b, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
