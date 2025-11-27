package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/stretchr/objx"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request by rendering the template.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]any{
		"Host": r.Host,
	}
	// If the user is authenticated, get the user data from the "auth" cookie
	if authcookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authcookie.Value)
	}

	t.templ.Execute(w, data)
}

func main() {
	// parse the command line flag for the address
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()
	// setup goth
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/auth/callback/google", "email", "profile"),
	)

	// create a new room
	r := newRoom(UseFileSystemAvatar)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/{action}/{provider}", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		// Invalidate the "auth" cookie by setting its MaxAge to -1
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload", MustAuth(&templateHandler{filename: "upload.html"}))
	http.Handle("/uploader", MustAuth(http.HandlerFunc(uploaderHandler)))
	// Serve avatar images from the "avatars" directory
	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))))
	// Get the room running
	// this will start the room's main event loop in the background to handle clients
	// joining, leaving and message forwarding, which allows the main goroutine to run the web server
	go r.run()
	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
