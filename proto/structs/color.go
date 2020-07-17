package structs

import (
	"github.com/tdakkota/go-terra/proto/common"
)

type Color struct {
	Red   byte
	Green byte
	Blue  byte
}

func (c Color) Len() int {
	return c.MinLength()
}

func (c Color) MinLength() int {
	return 3
}

func (c Color) Append(buf []byte) ([]byte, error) {
	var b []byte
	buf, b = common.Slice(buf, c.Len())

	b[0] = c.Red
	b[1] = c.Green
	b[2] = c.Blue

	return buf, nil
}

func (c Color) MarshalBinary() ([]byte, error) {
	return c.Append(make([]byte, 0, c.Len()))
}

func (c *Color) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.MinLength() {
		return common.ErrInvalidLength
	}

	c.Red = b[0]
	c.Green = b[1]
	c.Blue = b[2]
	return nil
}
