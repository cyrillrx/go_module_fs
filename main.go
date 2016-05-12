package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	path := path(os.Args)

	l := ls(path)

	prettyPrint(l)
}

func path(args []string) string {

	if len(args) > 1 {
		return args[1]
	} else {
		return "./"
	}
}

func ls(path string) []File {

	var output []File

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		output = append(output, makeFile(f))
	}

	return output
}

func prettyPrint(v interface{}) {

	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
