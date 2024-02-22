package main

import "github.com/gorilla/websocket"

// client represents a single chatting user
type client struct {
	// socket holds reference to websocket connection
	socket *websocket.Conn
	// room represents the room where client is in
	room *room
	// send is a channel on which messages are sent
	send chan []byte
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
