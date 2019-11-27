package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, user ID
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	if !alreadyLoggedIn(writer, request) {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	cookie, _ := request.Cookie("session")
	delete(dbSessions, cookie.Value)
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(writer, cookie)

	// clean up db sessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}

func index(writer http.ResponseWriter, request *http.Request) {
	u := getUser(writer, request)
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(writer, "index.gohtml", u)
}

func bar(writer http.ResponseWriter, request *http.Request) {
	u := getUser(writer, request)
	if !alreadyLoggedIn(writer, request) {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin" {
		http.Error(writer, "Permission denied, only admin can go to bar.", http.StatusForbidden)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(writer, "bar.gohtml", u)
}

func signup(writer http.ResponseWriter, request *http.Request) {
	if alreadyLoggedIn(writer, request) {
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
		r := request.FormValue("role")

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
		cookie.MaxAge = sessionLength
		http.SetCookie(writer, cookie)
		dbSessions[cookie.Value] = session{un, time.Now()}

		// store users in dbUsers
		bcp, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
		if err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
		}
		u := user{un, bcp, f, l, r}
		dbUsers[un] = u

		// redirect
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(writer, "signup.gohtml", nil)
}

func login(writer http.ResponseWriter, request *http.Request) {
	if alreadyLoggedIn(writer, request) {
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if request.Method == http.MethodPost {
		un := request.FormValue("username")
		pw := request.FormValue("password")
		// is there username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(writer, "Username and/or password do not match.", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pw))
		if err != nil {
			http.Error(writer, "Username and/or password do not match.", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		cookie.MaxAge = sessionLength
		http.SetCookie(writer, cookie)
		dbSessions[cookie.Value] = session{un, time.Now()}
		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(writer, "login.gohtml", nil)
}
