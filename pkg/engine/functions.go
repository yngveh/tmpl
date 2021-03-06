package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

var funcMap = template.FuncMap{
	"Env":       envFunc,
	"File":      fileFunc,
	"Date":      dateFunc,
	"Format":    formatFunc,
	"ToUpper":   strings.ToUpper,
	"ToLower":   strings.ToLower,
	"Title":     strings.Title,
	"Repeat":    strings.Repeat,
	"Replace":   strings.Replace,
	"Trim":      strings.Trim,
	"TrimSpace": strings.TrimSpace,
}

func envFunc(env string) string {
	return os.Getenv(env)
}

func fileFunc(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func dateFunc() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func formatFunc(text string, a ...interface{}) string {
	return fmt.Sprintf(text, a...)
}
