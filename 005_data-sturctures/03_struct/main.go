package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	gandalf := sage{
		Name:  "Gandalf the Grey",
		Motto: "Fly, you fools!",
	}

	err := tpl.Execute(os.Stdout, gandalf)
	if err != nil {
		log.Fatalln(err)
	}

}
