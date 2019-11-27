package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(writer http.ResponseWriter, request *http.Request) {
	// get cookie
	cookie, err := request.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(writer, cookie)
	}
	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if request.Method == http.MethodPost {
		un := request.FormValue("username")
		f := request.FormValue("firstname")
		l := request.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(writer, "index.gohtml", u)

}

func bar(writer http.ResponseWriter, request *http.Request) {

	// get cookie
	cookie, err := request.Cookie("session")
	if err != nil {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(writer, "bar.gohtml", u)

}
