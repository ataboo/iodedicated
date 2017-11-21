package wsserver

import (
	"github.com/gorilla/websocket"
	"log"
	"github.com/ataboo/iodedicated/game"
)

type User struct {
	conn *websocket.Conn
	hub game.EventHub
	player *game.Player
}

func NewUser (conn *websocket.Conn, player *game.Player) *User {
	user := User{conn, game.GetInstance().Hub, player}
	go user.ConnListener()
}

func (user User) ConnListener() {
	defer func() {
		user.conn.Close()
	}()
	for {
		_, message, err := user.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		user.hub.Incoming <- game.Event{user.player,message,}
	}
}

func (user User) Send(message string) {
	user.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (user User) Close() {
	user.conn.Close()
}