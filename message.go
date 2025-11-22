package main

import (
	"time"
)

// message represents a chat message.
type message struct {
	Name      string
	Message   string
	When      time.Time
	AvatarURL string
}
