package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
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
	if email, ok := c.userData["email"].(string); ok {
		// create MD5 hash of the email address, it calls whenever we need to get the avatar URL
		// need to improve the performance by caching the result
		m := md5.New()
		io.WriteString(m, strings.ToLower(strings.TrimSpace(email)))
		return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
	}
	return "", ErrNoAvatarURL
}
