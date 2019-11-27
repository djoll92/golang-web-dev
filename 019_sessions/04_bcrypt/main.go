package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
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
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	u := getUser(request)
	tpl.ExecuteTemplate(writer, "index.gohtml", u)
}

func bar(writer http.ResponseWriter, request *http.Request) {
	u := getUser(request)
	if !alreadyLoggedIn(request) {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(writer, "bar.gohtml", u)
}

func signup(writer http.ResponseWriter, request *http.Request) {
	if alreadyLoggedIn(request) {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if request.Method == http.MethodPost {

		// get form values
		un := request.FormValue("username")
		pw := request.FormValue("password")
		f := request.FormValue("firstname")
		l := request.FormValue("lastname")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(writer, "Username is already taken.", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(writer, cookie)
		dbSessions[cookie.Value] = un

		// store users in dbUsers
		bcp, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
		}
		u := user{un, bcp, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(writer, "signup.gohtml", nil)
}
