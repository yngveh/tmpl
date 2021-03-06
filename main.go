package main

import (
	"flag"
	"github.com/yngveh/tmpl/pkg/engine"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	tmplFlag := flag.String("tmpl", "", "Template filename. Use '-' to read from stdin.")
	dataFlag := flag.String("data", "", "Filename for file containing object data, json and yaml supported")
	flag.Parse()

	if *tmplFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	var (
		data   *map[interface{}]interface{}
		err    error
		source *os.File
	)

	if *dataFlag != "" {
		dataReader, err := os.Open(*dataFlag)
		checkError(err)
		defer dataReader.Close()

		data, err = engine.DataObject(dataReader)
		checkError(err)
	}

	if *tmplFlag == "-" {
		source = os.Stdin
	} else {
		source, err = os.Open(*tmplFlag)
		checkError(err)
	}
	defer source.Close()

	err = engine.Process(source, os.Stdout, data)
	checkError(err)

}
