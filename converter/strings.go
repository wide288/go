package converter

import "strconv"

func StrToByteArray(str string) []byte {
	return []byte(str)
}

func ByteArrayToStr(bA []byte) string {
	return string(bA[:])
}

func StrToUint64(str string) uint64 {
	uint64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint64
}

func StrToInt64(str string) int64 {
	int64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int64
}

func StrToInt(str string) int {
	intVal, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return intVal
}