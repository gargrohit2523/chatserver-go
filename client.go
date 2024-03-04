package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	// socket holds reference to websocket connection
	socket *websocket.Conn
	// room represents the room where client is in
	room *room
	// send is a channel on which messages are sent
	send chan *message
	// userData holds information about the user
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
