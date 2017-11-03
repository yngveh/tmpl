package main

import (
	"flag"
	"github.com/yngveh/tmpl/pkg/engine"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func dataObject(filename *string) (*map[interface{}]interface{}, error) {

	if *filename == "" {
		panic("Filen not found " + *filename)
	}

	d, err := ioutil.ReadFile(*filename)
	if err != nil {
		return nil, err
	}

	t := make(map[interface{}]interface{})
	err = yaml.Unmarshal(d, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func main() {

	tmplFlag := flag.String("tmpl", "", "File to read. To read from stdin use '-'")
	destFlag := flag.String("dest", "", "Destination file. If not present write to stdout")
	dataFlag := flag.String("data", "", "Filename for file containing object data, json and yaml supported")
	flag.Parse()

	if *tmplFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	var out io.Writer

	if *destFlag != "" {
		file, err := os.OpenFile(*destFlag, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		checkError(err)
		defer file.Close()
		out = file
	} else {
		out = os.Stdout
	}

	data, err := dataObject(dataFlag)
	if err != nil {
		panic(err)
	}

	err = engine.Process(tmplFlag, out, data)
	if err != nil {
		panic(err)
	}

}
