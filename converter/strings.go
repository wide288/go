package converter

func StrToByteArray(str string) []byte {
	return []byte(str)
}

func ByteArrayToStr(bA []byte) string {
	return string(bA[:])
}