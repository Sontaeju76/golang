package stjlibtcpcommon

// OBJMSG : message object
type OBJMSG interface {
	GetOPCode() uint16
	Parse(msg []byte) bool
	Make() []byte
}

/*
// OBJMSG : message
type OBJMSG struct {
	Header *OBJMSGHeader
	opCode uint16
}

// GetOPCode : get OPCode
func (obj *OBJMSG) GetOPCode() uint16 {
	rVal := obj.opCode
	if obj.Header != nil {
		rVal = obj.Header.OPCode
	}

	return rVal
}

// SetOPCode : set OPCode
func (obj *OBJMSG) SetOPCode(opcode uint16) {
	obj.opCode = opcode
}

// Parse : parsing message
func (obj *OBJMSG) Parse(msg []byte) bool {
	return false
}

// Make : make message
func (obj *OBJMSG) Make() []byte {
	return nil
}
*/
