package parser

import (
	"bytes"
	"io/ioutil"
	"os/exec"

	"github.com/yuin/goldmark"
)

// generateTemplate parses the source files into a template
func GenerateTemplate(layout, scss, markdown string, out *bytes.Buffer) {
	cmd := exec.Command("sass", scss)
	css, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	md, err := ioutil.ReadFile(markdown)
	if err != nil {
		panic(err)
	}
	var content bytes.Buffer
	if err = goldmark.Convert(md, &content); err != nil {
		panic(err)
	}

	html, err := ioutil.ReadFile(layout)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"css":     string(css),
		"content": content.String(),
	}
	ParseKeys("{%\x20(.*)\x20%}", html, data, out)
}
