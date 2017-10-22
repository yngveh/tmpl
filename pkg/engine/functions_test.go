package engine

import (
	"bytes"
	"os"
	"testing"
	"text/template"
)

func Test_should_process_env(t *testing.T) {
	os.Setenv("MY_ENV", "myenv")
	result := executeTemplate("{{ Env \"MY_ENV\" }}", t)
	assertString("myenv", result, t)
}

func Test_should_process_file(t *testing.T) {
	result := executeTemplate("{{ File \"../.././test/a_file.txt\" }}", t)
	assertString("sometext", result, t)
}

func Test_should_process_strings_functions(t *testing.T) {
	source := `
ToUpper       : {{ ToUpper "Text" }}
ToLower       : {{ ToLower "Text" }}
Repeat        : {{ Repeat "TEXT" 3 }}
Replace       : {{ Replace "TextText" "ex" "aa" 1 }}
Trim          : {{ Trim " ! Text ! " "! " }}
TrimSpace     : {{ TrimSpace " ! Text ! " }}
`
	expected := `
ToUpper       : TEXT
ToLower       : text
Repeat        : TEXTTEXTTEXT
Replace       : TaatText
Trim          : Text
TrimSpace     : ! Text !
`
	result := executeTemplate(source, t)
	assertString(expected, result, t)
}

func executeTemplate(source string, t *testing.T) string {
	t.Helper()
	tmpl, err := template.New("").Funcs(funcMap).Parse(source)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	tmpl.Execute(&b, nil)
	return b.String()
}

func assertString(expected, actual string, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
