package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func uploaderHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userid")
	file, header, err := r.FormFile("avatarFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// check if avatars directory exists, if not create it
	if _, err := os.Stat("avatars"); os.IsNotExist(err) {
		err = os.Mkdir("avatars", 0755)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	filename := path.Join("avatars", userId+path.Ext(header.Filename))
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Successful")
}
