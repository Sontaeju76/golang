package stjlibtcpserver

import (
	"net"
)

type netClient struct {
	Seq    int
	client net.Conn
}

// Init : initialize network client
func (c *netClient) Init(conn net.Conn) {
	c.client = conn

	go c.rcv()
}

// Send : send message
func (c *netClient) Send(msg []byte) {
	if c.client != nil {
		c.client.Write(msg)
	}
}

func (c *netClient) rcv() {
	buf := make([]byte, 1024)
	for {
		read, err := c.client.Read(buf)
		if err != nil {
			println("Client disconnected...", err)
			break
		}
		if read > 0 {
			println("RCV :: ", string(buf[:read]))
		}
	}
	// defer delete(Clients, c)
}
