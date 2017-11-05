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

	tmplFlag := flag.String("tmpl", "", "File to read. To read from stdin use '-'")
	dataFlag := flag.String("data", "", "Filename for file containing object data, json and yaml supported")
	flag.Parse()

	if *tmplFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	out := os.Stdout

	dataReader, err := os.Open(*dataFlag)
	checkError(err)
	defer dataReader.Close()

	data, err := engine.DataObject(dataReader)
	checkError(err)

	err = engine.Process(tmplFlag, out, data)
	checkError(err)

}
