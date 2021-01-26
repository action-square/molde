package mailer

import (
	"bytes"
	"net/smtp"
	"testing"

	"github.com/oagoulart/molde/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestMailer_Start(t *testing.T) {
	auth := smtp.PlainAuth("", "mail@localhost.com", "potato69", "localhost")
	data, err := parser.ReadJson("../../sample/data.json")
	assert.Nil(t, err)

	var template bytes.Buffer
	parser.GenerateTemplate("../../sample/layout.html",
		"../../sample/sass/styles.scss",
		"../../sample/content.md", &template)

	mailer := NewMailer(auth, "localhost:587", data, 4, &template)
	mailer.Start()
}
