package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurve adds CSRF protection to all POST request
func NoSurve(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session for every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}