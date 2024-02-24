package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (a authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")

	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	a.next.ServeHTTP(w, r)
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
