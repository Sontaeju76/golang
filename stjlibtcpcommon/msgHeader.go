package stjlibtcpcommon

import "time"

// OBJMSGHeader : message header
type OBJMSGHeader struct {
	PacketLength  uint32
	UserCode      uint32
	OPCode        uint16
	PacketID      uint32
	TotalPacket   uint16
	CurrentPacket uint16
	Timestamp     time.Time

	BodyLength uint32
	HasError   bool
}

// Parse :parse message header
func (h *OBJMSGHeader) Parse(header []byte) bool {
	var success bool
	offset := 4

	h.PacketLength, success = ByteArrToUInt32(&header, &offset, true)

	if success {
		h.UserCode, success = ByteArrToUInt32(&header, &offset, true)
	}
	if success {
		h.OPCode, success = ByteArrToUInt16(&header, &offset, true)
	}
	if success {
		h.PacketID, success = ByteArrToUInt32(&header, &offset, true)
	}
	if success {
		h.TotalPacket, success = ByteArrToUInt16(&header, &offset, true)
	}
	if success {
		h.CurrentPacket, success = ByteArrToUInt16(&header, &offset, true)
	}
	if success {
		h.Timestamp, success = ByteArrToTimeTimeval(&header, &offset, true)
	}

	h.HasError = !success

	return success
}

// Make :Make message header
func (h *OBJMSGHeader) Make() []byte {
	rVal := []byte{Preamble(0), Preamble(1), Preamble(2), 0x00}

	arrPacketLength := Uint32ToByteArr(h.PacketLength, true)
	rVal = append(rVal, arrPacketLength...)

	arrUserCode := Uint32ToByteArr(h.UserCode, true)
	rVal = append(rVal, arrUserCode...)

	arrOPCode := Uint16ToByteArr(h.OPCode, true)
	rVal = append(rVal, arrOPCode...)

	arrPacketID := Uint32ToByteArr(h.PacketID, true)
	rVal = append(rVal, arrPacketID...)

	arrTotalPacket := Uint16ToByteArr(h.TotalPacket, true)
	rVal = append(rVal, arrTotalPacket...)

	arrCurrentPacket := Uint16ToByteArr(h.CurrentPacket, true)
	rVal = append(rVal, arrCurrentPacket...)

	arrTimestamp := TimeToByteArrTimeval(h.Timestamp, true)
	rVal = append(rVal, arrTimestamp...)

	return rVal
}
