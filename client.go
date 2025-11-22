package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represents a connected user in the chat application.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan *message
	// room is the room this client is chatting in.
	room *room
	// userData holds information about the authenticated user.
	userData map[string]any
}

// read listens for incoming messages from the client's web socket
// and forwards them to the room's forward channel.
func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		name, ok := c.userData["name"].(string)
		if !ok {
			name = "unknown"
		}
		msg.Name = name
		if avatarURL, ok := c.userData["avatar_url"].(string); ok {
			msg.AvatarURL = avatarURL
		}
		// Forward the message to the room
		c.room.forward <- msg
	}
}

// write listens for messages on the send channel and writes them out to the client's web socket.
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
