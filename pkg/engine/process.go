package engine

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"text/template"
)

func Process(source io.Reader, out io.Writer, data *map[interface{}]interface{}) error {

	s, err := ioutil.ReadAll(source)
	if err != nil {
		return err
	}

	t, err := template.New("").Funcs(funcMap).Parse(string(s))
	if err != nil {
		return err
	}

	err = t.Execute(out, data)
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
