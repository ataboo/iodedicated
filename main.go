package main

import (
	"github.com/ataboo/iodedicated/wsserver"
)

func main() {
	server := wsserver.NewServer(":3000")

	server.Start()
}
