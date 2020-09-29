package stjlibtcpcommon

import (
	"encoding/binary"
	"math"
	"time"
)

// Uint16ToByteArr : uint16 to byte array
func Uint16ToByteArr(val uint16, isBigEndian bool) []byte {
	rVal := make([]byte, 2)

	if isBigEndian {
		binary.BigEndian.PutUint16(rVal, val)
	} else {
		binary.LittleEndian.PutUint16(rVal, val)
	}

	return rVal
}

// Uint32ToByteArr : uint32 to byte array
func Uint32ToByteArr(val uint32, isBigEndian bool) []byte {
	rVal := make([]byte, 4)

	if isBigEndian {
		binary.BigEndian.PutUint32(rVal, val)
	} else {
		binary.LittleEndian.PutUint32(rVal, val)
	}

	return rVal
}

// Uint64ToByteArr : uint64 to byte array
func Uint64ToByteArr(val uint64, isBigEndian bool) []byte {
	rVal := make([]byte, 8)

	if isBigEndian {
		binary.BigEndian.PutUint64(rVal, val)
	} else {
		binary.LittleEndian.PutUint64(rVal, val)
	}

	return rVal
}

// Int16ToByteArr : int16 to byte array
func Int16ToByteArr(val int16, isBigEndian bool) []byte {
	tmp := uint16(val)
	return Uint16ToByteArr(tmp, isBigEndian)
}

// Int32ToByteArr : int32 to byte array
func Int32ToByteArr(val int32, isBigEndian bool) []byte {
	tmp := uint32(val)
	return Uint32ToByteArr(tmp, isBigEndian)
}

// Int64ToByteArr : int64 to byte array
func Int64ToByteArr(val int64, isBigEndian bool) []byte {
	tmp := uint64(val)
	return Uint64ToByteArr(tmp, isBigEndian)
}

// Float32ToByteArr : float32 to byte array
func Float32ToByteArr(val float32, isBigEndian bool) []byte {
	rVal := make([]byte, 4)
	bits := math.Float32bits(val)
	if isBigEndian {
		binary.BigEndian.PutUint32(rVal, bits)
	} else {
		binary.LittleEndian.PutUint32(rVal, bits)
	}

	return rVal
}

// Float64ToByteArr : float64 to byte array
func Float64ToByteArr(val float64, isBigEndian bool) []byte {
	rVal := make([]byte, 8)
	bits := math.Float64bits(val)
	if isBigEndian {
		binary.BigEndian.PutUint64(rVal, bits)
	} else {
		binary.LittleEndian.PutUint64(rVal, bits)
	}

	return rVal
}

// TimeToByteArrTimeval : Time to byte array. 8bytes c timeval
func TimeToByteArrTimeval(val time.Time, isBigEndian bool) []byte {
	rVal := make([]byte, 8)
	sec := val.Unix()
	msec := (val.UnixNano() / 1000) % 1000000

	tmp1 := Uint32ToByteArr(uint32(sec), isBigEndian)
	tmp2 := Uint32ToByteArr(uint32(msec), isBigEndian)

	idx := 0
	for _, i := range tmp1 {
		rVal[idx] = i
		idx++
	}
	for _, i := range tmp2 {
		rVal[idx] = i
		idx++
	}

	return rVal
}

// TimeToByteArrUnixtimestamp : Time to byte array. 4bytes unix Timestamp
func TimeToByteArrUnixtimestamp(val time.Time, isBigEndian bool) []byte {
	sec := val.Unix()
	rVal := Uint32ToByteArr(uint32(sec), isBigEndian)
	return rVal
}

// StringToByteArr : string to byte array.
func StringToByteArr(val string, length int) []byte {
	var rVal []byte
	tmp := []byte(val)

	tmpLen := len(tmp)

	if length == tmpLen {
		rVal = tmp
	} else {
		rVal = make([]byte, length, length)
		copy(rVal, tmp)
	}

	return rVal
}
