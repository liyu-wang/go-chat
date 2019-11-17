package main

import (
	"errors"
	"io/ioutil"
	"path"

	gomniauthcommon "github.com/stretchr/gomniauth/common"
)

// ChatUser interface
type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

// ErrNoAvatarURL is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

// Avatar represents types capable of representing
// user profile pictures.
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNOAvatarURL is returned if the object is unable to get
	// a URL for the specified client.
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars is a slice of available Avatar implementations
type TryAvatars []Avatar

// GetAvatarURL try all the available implementations
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar is a concrete class that implements Avatar
type AuthAvatar struct{}

// UseAuthAvatar nil variable
var UseAuthAvatar AuthAvatar

// GetAvatarURL implementation of AuthAvatar
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return url, nil
}

// GravatarAvatar is a concrete class that implements Avatar
type GravatarAvatar struct{}

// UseGravatar nil variable
var UseGravatar GravatarAvatar

// GetAvatarURL implementation of GravatarAvatar
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSytemAvatar is a concrete class that implements Avatar
type FileSytemAvatar struct{}

// UseFileSystemAvatar nil variable
var UseFileSystemAvatar FileSytemAvatar

// GetAvatarURL implementation of FileSytemAvatar
func (FileSytemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
