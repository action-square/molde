package main

import (
	"bytes"
	"errors"
	"flag"
	"log"
	"net/smtp"
	"os"

	"github.com/oagoulart/molde/pkg/mailer"
	"github.com/oagoulart/molde/pkg/parser"
)

func main() {
	var authId, authUser, authPass string
	flag.StringVar(&authId, "authId", "", "`optional` usually not required for authentication")
	flag.StringVar(&authUser, "authUser", "", "`username` for authentication")
	flag.StringVar(&authPass, "authPass", "", "`password` for authentication")
	var from, host, port string
	flag.StringVar(&from, "from", "", "sender mail username")
	flag.StringVar(&host, "host", "", "mail server `address`")
	flag.StringVar(&port, "port", "587", "mail server `port`")
	var data, layout, styles, content string
	flag.StringVar(&data, "data", "./data.json", "mail `data` JSON file")
	flag.StringVar(&layout, "layout", "./layout.html", "mail `layout` HTML file")
	flag.StringVar(&styles, "styles", "./sass/styles.scss", "mail `styles` Sass/SCSS main file")
	flag.StringVar(&content, "content", "./content.md", "mail `content` Markdown file")
	var numWorkers uint
	flag.UintVar(&numWorkers, "workers", 4, "number of parallel `workers` at the same time")
	var showHelp bool
	flag.BoolVar(&showHelp, "help", false, "show help message")
	flag.Parse()

	if showHelp {
		flag.PrintDefaults()
		os.Exit(1)
	} else if host == "" {
		log.Fatal(errors.New("you need to specify --host. use --help for usage"))
	}

	auth := smtp.PlainAuth(authId, authUser, authPass, host)
	parsedData, err := parser.ReadJson(data)
	if err != nil {
		log.Fatal(err)
	}

	var template bytes.Buffer
	parser.GenerateTemplate(layout, styles, content, &template)

	mailer := mailer.NewMailer(auth, host+":"+port, from, parsedData, numWorkers, &template)
	mailer.Start()
}
