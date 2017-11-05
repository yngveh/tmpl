package engine

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
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

func DataObject(dataReader io.Reader) (*map[interface{}]interface{}, error) {

	d, err := ioutil.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}

	j := make(map[string]interface{})
	err = json.Unmarshal(d, &j)
	if err == nil {
		d, _ = yaml.Marshal(&j)
	}

	var t map[interface{}]interface{}
	err = yaml.Unmarshal(d, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
