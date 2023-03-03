package main

import "github.com/gameloungee/client/internal/server"

func main() {
	server := new(server.Server)
	addr := server.MakeAddr()
	server.Run(addr)
}
