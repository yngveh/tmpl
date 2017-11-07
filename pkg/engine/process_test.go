package engine

import (
	"bytes"
	"strings"
	"testing"
)

func Test_should_process(t *testing.T) {

	data := make(map[interface{}]interface{})
	data["key"] = "value"

	out := bytes.NewBufferString("")
	template := bytes.NewBufferString("{{ .key }}")

	err := Process(template, out, &data)

	if err != nil {
		t.Errorf("Failed parsing template %v, with: %v", template.String(), err)
	}

	if out.String() != "value" {
		t.Errorf("Expected: value, Actual: '%s'", out.String())
	}
}

func Test_source_convert_yaml_to_data_object(t *testing.T) {

	data := `
root:
  var1: value1
  var2: value2
`
	input := strings.NewReader(data)

	result, err := DataObject(input)
	if err != nil {
		t.Errorf("Failed with: %v", err)
	}

	mapObject := *result
	root := mapObject["root"]
	if root == nil {
		t.Errorf("Failed converting: %v", root)
	}

}

func Test_source_convert_json_to_data_object(t *testing.T) {

	data := `
{
	"root":{
		"var1":"value1",
  	  	"var2":"value2"
	}
}
`
	input := strings.NewReader(data)

	result, err := DataObject(input)
	if err != nil {
		t.Errorf("Failed with: %v", err)
	}

	mapObject := *result
	root := mapObject["root"]
	if root == nil {
		t.Errorf("Failed converting: %v", root)
	}

}
