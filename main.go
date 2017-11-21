package main

import (
	"github.com/ataboo/iodedicated/wsserver"
	"github.com/ataboo/iodedicated/game"
)

func main() {
	server := wsserver.NewServer(":3000")

	server.Start()
	game.GetInstance()
}
