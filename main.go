package main

import (
	"flag"
	"io"
	"os"
	"path"
	"text/template"
)

var funcMap = template.FuncMap{
	"Env": envFunc,
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func envFunc(env string) string {
	return os.Getenv(env)
}

func process(filename *string, out io.Writer) {
	t, err := template.New("").Funcs(funcMap).ParseFiles(*filename)
	checkError(err)
	_, tmplName := path.Split(*filename)
	err = t.ExecuteTemplate(out, tmplName, nil)
	checkError(err)
}

func main() {

	tmplFlag := flag.String("tmpl", "", "File to read. To read from stdin use '-'")
	destFlag := flag.String("dest", "", "Destination file. If not present write to stdout")
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

	process(tmplFlag, out)

}
