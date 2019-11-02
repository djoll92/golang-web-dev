package main

import (
	"log"
	"os"
	"strings"
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

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
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

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
