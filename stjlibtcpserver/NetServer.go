package stjlibtcpserver

import (
	"net"
	"strconv"
)

// Clients : connected clients
var Clients map[*netClient]int

func init() {
	Clients = make(map[*netClient]int)
}

// Start : start server
func Start(port int) {

	if port > 0 {
		port := ":" + strconv.Itoa(port)

		go open(port)
	}
}

func open(port string) {
	conn, err := net.Listen("tcp", port)
	if err != nil {
		println(err)
	}
	defer conn.Close()

	for {
		c, err := conn.Accept()
		if err != nil {
			println(err)
		}

		var client netClient
		client.Init(c)
		Clients[&client] = 1
	}
}
