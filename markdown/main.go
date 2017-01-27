package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title   string
	Content string
}

// markdown author is trusted
func html(file []byte) {
	content := blackfriday.MarkdownCommon(file)
	page := Page{"Markdown sample", string(content)}

	tmpl, err := template.ParseFiles("template/template.tmpl")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	var fd = os.Stdin
	var err error
	var file []byte

	if len(os.Args) > 1 {
		fd, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	file, err = ioutil.ReadAll(fd)
	err = fd.Close()

	html(file)
}
