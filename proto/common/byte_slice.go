package common

import "math"

const (
	MaxBytesLength = math.MaxUint8
	MinBytesLength = 0
)

func ReadBytes(b []byte) ([]byte, error) {
	return ReadBytesRange(b, MinBytesLength, MaxBytesLength)
}

func ReadBytesRange(b []byte, min, max int) ([]byte, error) {
	length := int(b[0])
	switch length {
	case 0:
		return []byte{}, nil
	case 1:
		return []byte{b[1]}, nil
	}

	if len(b)-1 < length {
		return nil, ErrInvalidLength
	}

	if length < min || length > max {
		return nil, LengthConstraintError{min, max}
	}

	return b[1 : length+1], nil
}

func WriteBytesRange(s []byte, b []byte, min, max int) error {
	length := len(s)
	if length > math.MaxUint8 {
		return ErrStringTooLong
	}

	if length <= min || length > max {
		return LengthConstraintError{min, max}
	}

	b[0] = byte(len(s))
	copy(b[1:], s)
	return nil
}

func WriteBytes(s []byte, b []byte) error {
	return WriteBytesRange(s, b, MinBytesLength, MaxBytesLength)
}
