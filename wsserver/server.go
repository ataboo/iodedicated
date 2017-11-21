package wsserver

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"github.com/ataboo/iodedicated/game"
)

type WsHost struct {
	Address string
	Server *http.Server
	Upgrader websocket.Upgrader
}

func NewServer(address string) *WsHost {
	host := &WsHost{
		Address: address,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func (r *http.Request) bool {
				fmt.Println(r)
				return true
			},
		},
	}

	return host
}

func (host *WsHost) Start() {
	host.Server = &http.Server{ Addr: host.Address }
	http.HandleFunc("/ws", host.handleWs)
	host.Server.ListenAndServe()
}

func (host *WsHost) Stop() {
	host.Server.Close()
}

func (host *WsHost) handleWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tadaa!")

	conn, err := host.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade err: ", err)
		return
	}

	userName := r.URL.Query().Get("username")

	fmt.Println("Username: "+userName+" connected.")

	player := game.GetInstance().Roster.FindOrNew(userName)
	player.SetUser(*NewUser(conn, player))
}
