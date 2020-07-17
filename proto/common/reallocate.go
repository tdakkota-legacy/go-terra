package common

func AppendSize(b []byte, length int) (r []byte) {
	if cap(b)-len(b) < length {
		r = make([]byte, len(b)+length)
		copy(r, b)
		return r
	}

	return b
}

func Slice(b []byte, length int) (old, write []byte) {
	oldLen := len(b)
	b = AppendSize(b, length)
	return b[:length+oldLen], b[oldLen : length+oldLen]
}
