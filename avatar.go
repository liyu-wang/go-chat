package main

import (
	"errors"
	"os"
	"path"
)

// ErrNoAvatarURL is returned when the Avatar instance is unable to provide
// an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get avatar URL")

// Avatar represents types capable of repreenting user proifile pictures.
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client.
	// or returns ErrNoAvatarURL if the avatar URL is not available.
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"].(string); ok {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"].(string); ok {
		return "//www.gravatar.com/avatar/" + userid, nil
	}
	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"].(string); ok {
		files, err := os.ReadDir("avatars")
		if err != nil {
			return "", ErrNoAvatarURL
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(userid+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
