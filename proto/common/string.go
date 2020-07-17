package common

import (
	"math"
)

const (
	MaxStringLength = math.MaxUint8
	MinStringLength = 0
)

func ReadString(b []byte) (string, error) {
	return ReadStringRange(b, MinStringLength, MaxStringLength)
}

func ReadStringRange(b []byte, min, max int) (string, error) {
	length := int(b[0])
	switch length {
	case 0:
		return "", nil
	case 1:
		return string(b[1]), nil
	}

	if len(b)-1 < length {
		return "", ErrInvalidLength
	}

	if length < min || length > max {
		return "", LengthConstraintError{min, max}
	}

	return string(b[1 : length+1]), nil
}

func WriteStringRange(s string, b []byte, min, max int) error {
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

func WriteString(s string, b []byte) error {
	return WriteStringRange(s, b, MinStringLength, MaxStringLength)
}
