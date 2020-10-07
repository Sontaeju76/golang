package stjlibtcpserver

import (
	"net"
	netcomm "stj/stjlibtcpcommon"
	"strconv"
)

// Clients : connected clients
// var Clients map[*netClient]int
var Clients map[uint32]*netcomm.NetClient
var clientsSeq uint32

var rcvch chan *netcomm.OBJMSGARGS

func init() {
	// Clients = make(map[*netClient]int)
	Clients = make(map[uint32]*netcomm.NetClient)
	clientsSeq = 1
	rcvch = make(chan *netcomm.OBJMSGARGS, 200)
}

// Start : start server
func Start(port int, ch chan *netcomm.OBJMSGARGS) {

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

		// var client netClient
		// Clients[&client] = 1
		client := netcomm.NetClient{Seq: clientsSeq, Client: c, Ch: rcvch}
		client.Init()

		Clients[clientsSeq] = &client
		clientsSeq++
	}
}

func rcvProc(ch chan *netcomm.OBJMSGARGS) {
	for {
		msg := <-rcvch
		if msg.Header == nil {
			delete(Clients, msg.ClientSeq)
		}

		ch <- msg
	}
}
