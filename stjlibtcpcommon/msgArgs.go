package stjlibtcpcommon

// HeaderSize : message header size
const HeaderSize int = 30

// TailSize : message tail size
const TailSize int = 4

// PREAMBLE : message preamble
var PREAMBLE = [3]byte{0x45, 0x4D, 0x50}

// Preamble : message Preamble
// func Preamble(idx int8) byte {
// 	if idx == 0 {
// 		return 0x45
// 	} else if idx == 1 {
// 		return 0x4D
// 	} else if idx == 2 {
// 		return 0x50
// 	}
// 	return 0x00
// }

func getPacketID() func() uint {
	var i uint
	return func() uint {
		i++
		return i
	}
}

// PacketID : send packet ID
var PacketID = getPacketID()

// OBJMSGARGS : message arguments
type OBJMSGARGS struct {
	ClientSeq uint32
	Header    *OBJMSGHeader
	Data      []byte
}
