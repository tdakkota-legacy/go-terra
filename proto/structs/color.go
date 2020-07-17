package structs

import (
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type Color struct {
	Red   byte
	Green byte
	Blue  byte
}

func (c Color) minLength() int {
	return 3
}

func (c Color) MarshalBinary() ([]byte, error) {
	n := [3]byte{c.Red, c.Green, c.Blue}
	return n[:], nil
}

func (c *Color) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.minLength() {
		return common2.ErrInvalidLength
	}

	c.Red = b[0]
	c.Green = b[1]
	c.Blue = b[2]
	return nil
}
