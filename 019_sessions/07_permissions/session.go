package main

import "net/http"

func getUser(r *http.Request) user {
	var u user

	// get cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get the user
	if un, ok := dbSessions[cookie.Value]; ok {
		return dbUsers[un]
	}

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok

}
