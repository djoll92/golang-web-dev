package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, r *http.Request) user {

	// get cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	var u user

	// if the user exists already, get the user
	if s, ok := dbSessions[cookie.Value]; ok {
		return dbUsers[s.un]
	}

	return u
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}

	s, ok := dbSessions[cookie.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
	}

	_, ok = dbUsers[s.un]
	// refresh session
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)
	return ok

}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
