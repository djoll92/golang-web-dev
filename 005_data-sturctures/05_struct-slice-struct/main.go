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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	gandalf := sage{
		Name:  "Gandalf the Grey",
		Motto: "Fly, you fools!",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	ford := car{
		Manufacturer: "Ford",
		Model:        "Escort",
		Doors:        5,
	}

	toyota := car{
		Manufacturer: "Toyota",
		Model:        "Yaris",
		Doors:        4,
	}

	sages := []sage{gandalf, gandhi, buddha}
	cars := []car{ford, toyota}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
