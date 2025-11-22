package main

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/stretchr/objx"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	} else if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// success - call the next handler
	h.next.ServeHTTP(w, r)
}

// MustAuth helper function simply creates authHandler that wraps any other http.Handler
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

// loginHandler handles the third-party login process.
// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	action := r.PathValue("action")
	provider := r.PathValue("provider")

	switch action {
	case "login":
		// try to get the user without re-authenticating
		// gothic.getProviderName automatically checks the provider path value
		if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			fmt.Fprintf(w, "User already logged in: %v", gothUser)
			return
		} else {
			gothic.BeginAuthHandler(w, r)
		}

	case "callback":
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintf(w, "Error completing authentication: %s", err)
			http.Error(w, fmt.Sprintf("Error when trying to complete auth from %s: %s", provider, err), http.StatusInternalServerError)
			return
		}

		// create the auth cookie value
		authCookieValue := objx.New(map[string]any{
			"name": user.Name,
		}).MustBase64()
		// set the auth cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
