package wsserver

import (
	"github.com/gorilla/websocket"
)

type Player struct{
	id int
	conn websocket.Conn
}

func (player Player) Send(message string) {
	player.conn.WriteMessage(websocket.TextMessage, []byte(message))
}


