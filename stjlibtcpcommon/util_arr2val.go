package stjlibtcpcommon

import (
	"encoding/binary"
	"math"
	"time"
)

// ByteArrToUInt16 : byte array to uint16
func ByteArrToUInt16(arr *[]byte, offset *int, isBigEndian bool) (uint16, bool) {
	var rVal uint16
	var success bool

	ridx := (*offset) + 2

	if ridx <= len(*arr) {
		tmp := (*arr)[*offset:ridx]

		if isBigEndian {
			rVal = binary.BigEndian.Uint16(tmp)
		} else {
			rVal = binary.LittleEndian.Uint16(tmp)
		}

		*offset = ridx
		success = true
	}

	return rVal, success
}

// ByteArrToUInt32 : byte array to uint32
func ByteArrToUInt32(arr *[]byte, offset *int, isBigEndian bool) (uint32, bool) {
	var rVal uint32
	var success bool

	ridx := (*offset) + 4

	if ridx <= len(*arr) {
		tmp := (*arr)[*offset:ridx]

		if isBigEndian {
			rVal = binary.BigEndian.Uint32(tmp)
		} else {
			rVal = binary.LittleEndian.Uint32(tmp)
		}

		*offset = ridx
		success = true
	}

	return rVal, success
}

// ByteArrToUInt64 : byte array to uint64
func ByteArrToUInt64(arr *[]byte, offset *int, isBigEndian bool) (uint64, bool) {
	var rVal uint64
	var success bool

	ridx := (*offset) + 8

	if ridx <= len(*arr) {
		tmp := (*arr)[*offset:ridx]

		if isBigEndian {
			rVal = binary.BigEndian.Uint64(tmp)
		} else {
			rVal = binary.LittleEndian.Uint64(tmp)
		}

		*offset = ridx
		success = true
	}

	return rVal, success
}

// ByteArrToInt16 : byte array to int16
func ByteArrToInt16(arr *[]byte, offset *int, isBigEndian bool) (int16, bool) {
	unsignedNum, success := ByteArrToUInt16(arr, offset, isBigEndian)
	return int16(unsignedNum), success
}

// ByteArrToInt32 : byte array to int32
func ByteArrToInt32(arr *[]byte, offset *int, isBigEndian bool) (int32, bool) {
	unsignedNum, success := ByteArrToUInt32(arr, offset, isBigEndian)
	return int32(unsignedNum), success
}

// ByteArrToInt64 : byte array to int64
func ByteArrToInt64(arr *[]byte, offset *int, isBigEndian bool) (int64, bool) {
	unsignedNum, success := ByteArrToUInt64(arr, offset, isBigEndian)
	return int64(unsignedNum), success
}

// ByteArrToFloat32 : byte array to float32
func ByteArrToFloat32(arr *[]byte, offset *int, isBigEndian bool) (float32, bool) {
	tmpNum, success := ByteArrToUInt32(arr, offset, isBigEndian)
	return math.Float32frombits(tmpNum), success
}

// ByteArrToFloat64 : byte array to float64
func ByteArrToFloat64(arr *[]byte, offset *int, isBigEndian bool) (float64, bool) {
	tmpNum, success := ByteArrToUInt64(arr, offset, isBigEndian)
	return math.Float64frombits(tmpNum), success
}

// ByteArrToTimeTimeval : byte array to Time. 8bytes c timeval
func ByteArrToTimeTimeval(arr *[]byte, offset *int, isBigEndian bool) (time.Time, bool) {
	var rVal time.Time
	tmpNum1, success := ByteArrToUInt32(arr, offset, isBigEndian)
	if success {
		tmpNum2, success := ByteArrToUInt32(arr, offset, isBigEndian)

		if success {
			rVal = time.Unix(int64(tmpNum1), int64(tmpNum2))
		}
	}

	return rVal, success
}

// ByteArrToTimeUnixtimestamp : byte array to Time. 4bytes unix Timestamp
func ByteArrToTimeUnixtimestamp(arr *[]byte, offset *int, isBigEndian bool) (time.Time, bool) {
	var rVal time.Time
	tmpNum1, success := ByteArrToUInt32(arr, offset, isBigEndian)
	if success {
		rVal = time.Unix(int64(tmpNum1), 0)
	}

	return rVal, success
}

// ByteArrToString : byte array to string. utf8 & ascii
func ByteArrToString(arr *[]byte, offset *int, length int) (string, bool) {
	var rVal string
	var success bool

	ridx := (*offset) + length

	if ridx <= len(*arr) {
		tmp := (*arr)[*offset:ridx]

		rVal = string(tmp)
		success = true
	}

	return rVal, success
}
