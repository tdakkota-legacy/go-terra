package common

func WriteBool(b bool, buf []byte) {
	if b {
		buf[0] = 1
	} else {
		buf[0] = 0
	}
}

func ReadBool(buf []byte) bool {
	return buf[0] > 0
}
