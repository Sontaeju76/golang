package stjlibtcpcommon

import "net"

// NetClient : network client
type NetClient struct {
	Seq    int
	Client net.Conn
	Ch     chan<- OBJMSGARGS
}

// Init : initialize network client
func (c *NetClient) Init(conn net.Conn) {
	c.Client = conn

	go c.rcv()
}

// Send : send message
func (c *NetClient) Send(msg []byte) {
	if c.Client != nil {
		c.Client.Write(msg)
	}
}

func (c *NetClient) rcv() {
	buf := make([]byte, 1024)
	for {
		read, err := c.Client.Read(buf)
		if err != nil {
			println("Client disconnected...", err)
			break
		}
		if read > 0 {
			println("RCV :: ", string(buf[:read]))

			var msgarg OBJMSGARGS
			c.Ch <- msgarg
		}
	}
	//	defer delete(Clients, c)
}
