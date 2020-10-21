package stjlibtcpcommon

import "net"

// BUFFSIZE : buffer size
const BUFFSIZE int = 102400

// NetClient : network client
type NetClient struct {
	Seq    uint32
	Client net.Conn
	Ch     chan<- *OBJMSGARGS

	NeedHeader    bool
	MessageBuffer [BUFFSIZE]byte
	BuffIdxS      int
	BuffIdxE      int
	BuffDataSize  int
}

func (c *NetClient) getData(length int) []byte {
	var rVal []byte

	return rVal
}

func (c *NetClient) atData(idx int) byte {
	var rVal byte

	tmpIdx := c.BuffIdxS + idx
	if tmpIdx >= BUFFSIZE {
		tmpIdx = tmpIdx - BUFFSIZE
	}

	rVal = c.MessageBuffer[tmpIdx]

	return rVal
}

func (c *NetClient) moveIdxS(leng int) {
	if c.BuffDataSize > leng {
		tmpIdxS := c.BuffIdxS + leng

		if tmpIdxS >= BUFFSIZE {
			tmpIdxS -= BUFFSIZE
		}

		if tmpIdxS < c.BuffIdxE {
			c.BuffIdxS = tmpIdxS
			c.BuffDataSize = c.BuffIdxE - c.BuffIdxS
		} else {
			c.BuffIdxS = 0
			c.BuffIdxE = 0
			c.BuffDataSize = 0
		}
	} else {
		c.BuffIdxS = 0
		c.BuffIdxE = 0
		c.BuffDataSize = 0
	}
}

func (c *NetClient) schPreamble() bool {
	var rVal bool

	if c.MessageBuffer[c.BuffIdxS] == PREAMBLE[0] && c.MessageBuffer[c.BuffIdxS+1] == PREAMBLE[1] && c.MessageBuffer[c.BuffIdxS+2] == PREAMBLE[2] {
		rVal = true
	} else {
		idxE := c.BuffDataSize - 2
		if idxE > 1 {
			for i := 1; i < idxE; i++ {
				if c.atData(i) == PREAMBLE[0] {
					if c.atData(i+1) == PREAMBLE[1] && c.atData(i+2) == PREAMBLE[2] {
						c.moveIdxS(i)
						rVal = true
						break
					}
				}
			}

			if !rVal {
				c.moveIdxS(c.BuffDataSize - 1)
			}
		}
	}

	return rVal
}

// Init : initialize network client
func (c *NetClient) Init() {
	if c.Client != nil {
		c.NeedHeader = true
		go c.rcv()
	}
}

// AppendBuffer : append buffer
func (c *NetClient) AppendBuffer(data []byte) {
	leng := len(data)
	c.BuffDataSize += leng

	tmpIdxE := c.BuffIdxE + leng
	if tmpIdxE < BUFFSIZE {
		for i, b := range data {
			c.MessageBuffer[c.BuffIdxE+i] = b
		}
		c.BuffIdxE = tmpIdxE
	} else {
		tmpIdxE = tmpIdxE - BUFFSIZE
		var j int
		for i := c.BuffIdxE; i < BUFFSIZE; i++ {
			c.MessageBuffer[i] = data[j]
			j++
		}
		for i := 0; i < tmpIdxE; i++ {
			c.MessageBuffer[i] = data[j]
			j++
		}
		c.BuffIdxE = tmpIdxE
	}
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
			c.AppendBuffer(buf[:read])

			if c.BuffDataSize >= HeaderSize {
				if c.schPreamble() {
				}
			}

			if !c.NeedHeader {
			}

			println("RCV :: ", string(buf[:read]))

			var msgarg OBJMSGARGS
			c.Ch <- &msgarg
		}
	}

	var msgarg OBJMSGARGS
	c.Ch <- &msgarg

	//	defer delete(Clients, c)
}
