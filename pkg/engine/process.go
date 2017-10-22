package engine

import (
	"io"
	"path"
	"text/template"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Process(filename *string, out io.Writer) {
	t, err := template.New("").Funcs(funcMap).ParseFiles(*filename)
	checkError(err)
	_, tmplName := path.Split(*filename)
	err = t.ExecuteTemplate(out, tmplName, nil)
	checkError(err)
}
