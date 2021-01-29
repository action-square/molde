package parser

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParseKeys(t *testing.T) {
	template := `
    <html>
      <head></head>
      <body>
        <h1>{{ title }}</h1>
        <p>{{ body }}</p>
      </body>
    </html>
  `

	keys := map[string]interface{}{"title": "potato", "body": "banana"}

	var output bytes.Buffer
	ParseKeys("{{\x20(.*)\x20}}", []byte(template), keys, &output)

	expect := `
    <html>
      <head></head>
      <body>
        <h1>potato</h1>
        <p>banana</p>
      </body>
    </html>
  `
	assert.Equal(t, expect, output.String(), "The keys were not parsed correctly.")
}

func TestParser_ReadJson(t *testing.T) {
	out, err := ReadJson("../../sample/data.json")
	assert.Nil(t, err)

	expect := map[string]interface{}{
		"to":      "me@localhost.com",
		"subject": "your potato",
		"name":    "poopoopicker"}
	assert.Equal(t, expect, out[0], "The output map is not correct.")
}
