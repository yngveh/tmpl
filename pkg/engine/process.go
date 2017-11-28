package engine

import (
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

	if err = t.Execute(out, data); err != nil {
		return err
	}
	return nil
}

func DataObject(dataReader io.Reader) (*map[interface{}]interface{}, error) {

	data, err := ioutil.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}

	var yamlMap map[interface{}]interface{}
	err = yaml.Unmarshal(data, &yamlMap)
	if err != nil {
		return nil, err
	}

	return &yamlMap, nil
}
