package engine

import (
	"bytes"
	"testing"
)

func Test_should_process(t *testing.T) {

	data := make(map[interface{}]interface{})
	data["key"] = "value"

	b := bytes.NewBufferString("")
	filename := "../../test/data.tpl"
	err := Process(&filename, b, &data)

	if err != nil {
		t.Errorf("Failed parsing template %s, with: %v", filename, err)
	}

	if b.String() != "value" {
		t.Errorf("Expected: value, Actual: '%s'", b.String())
	}
}
