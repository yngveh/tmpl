package engine

import (
	"io"
	"path"
	"text/template"
)

func Process(filename *string, out io.Writer, data *map[interface{}]interface{}) error {
	t, err := template.New("").Funcs(funcMap).ParseFiles(*filename)
	if err != nil {
		return err
	}

	_, tmplName := path.Split(*filename)
	err = t.ExecuteTemplate(out, tmplName, data)
	if err != nil {
		return err
	}
	return nil
}
