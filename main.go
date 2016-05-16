package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Handles exe flags
	ppArg := flag.String("p", "inline", "true to pretty print, false to inline")
	//json := *flag.String("f", "json", "json or raw") == "json"

	flag.Parse()
	args := flag.Args()

	path := path(args)

	// List files
	files := ls(path)

	pp := *ppArg == "pretty"
	jsonPrint(files, pp)
}

func path(args []string) string {

	if len(args) > 0 {
		return args[0]
	} else {
		return "./"
	}
}

func ls(path string) []File {

	var output []File

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	for _, f := range files {
		output = append(output, makeFile(f))
	}

	return output
}

func jsonPrint(v interface{}, pretty bool) {

	b, err := serialize(v, pretty)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func serialize(v interface{}, pretty bool) ([]byte, error) {

	if pretty {
		return json.MarshalIndent(v, "", "    ")
	} else {
		return json.Marshal(v)
	}
}
