package mailer

import (
	"bytes"
	"log"
	"net/smtp"

	"github.com/oagoulart/molde/pkg/parser"
)

// Mailer is used to store mailing information
type Mailer struct {
	auth       smtp.Auth
	address    string
	data       []interface{}
	numWorkers uint
	template   *bytes.Buffer
}

// NewMailer creates new `Mailer` instance
func NewMailer(auth smtp.Auth, address string, data []interface{}, numWorkers uint, template *bytes.Buffer) *Mailer {
	return &Mailer{
		auth:       auth,
		address:    address,
		data:       data,
		numWorkers: numWorkers,
		template:   template,
	}
}

// mail is used to create a new mailer routine
func (m *Mailer) mail(keys map[string]interface{}, template []byte, sig <-chan bool, code chan<- uint, errs chan<- error) {
	<-sig

	var body bytes.Buffer
	parser.ParseKeys("{{\x20(.*)\x20}}", template, keys, &body)

	head := "To: " + keys["to"].(string) + "\r\n" +
		"Subject: " + keys["subject"].(string) + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n" +
		"Content-Transfer-Encoding: 7bit\r\n\r\n"
	msg := []byte(head + "\n" + body.String() + "\r\n")
	to := []string{keys["to"].(string)}

	if err := smtp.SendMail(m.address, m.auth, keys["from"].(string), to, msg); err != nil {
		errs <- err
	} else {
		code <- 400
	}
}

// Start initializes the mailing job
func (m *Mailer) Start() {
	sig := make(chan bool, m.numWorkers)
	code := make(chan uint)
	errs := make(chan error)
	defer close(sig)

	for _, mail := range m.data {
		go m.mail(mail.(map[string]interface{}), m.template.Bytes(), sig, code, errs)
	}

	for i := 0; uint(i) < m.numWorkers; i++ {
		sig <- true
	}

	toSend := len(m.data)
	for toSend > 0 {
		select {
		case <-code:
			toSend--
			sig <- true
		case e := <-errs:
			log.Println(e.Error())

			toSend--
			sig <- true
		}
	}
}
