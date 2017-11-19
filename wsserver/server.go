package wsserver

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
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

	c, err := host.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade err: ", err)
		return
	}

	fmt.Println(c)

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			break
		}
		log.Printf("received: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write err: ", err)
			break
		}
	}
}
