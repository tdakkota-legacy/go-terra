package structs

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

type Chest struct {
	Index int16
	X     int16
	Y     int16
	Name  string
}

func (c Chest) Len() int {
	return c.MinLength() + len(c.Name)
}

func (c Chest) MinLength() int {
	return 3*2 + 1 // Index(2) + X(2) + Y(2) + Len(1)
}

func (c Chest) Append(buf []byte) ([]byte, error) {
	var b []byte
	buf, b = common.Slice(buf, c.Len())

	binary.LittleEndian.PutUint16(b[0:], uint16(c.Index))
	binary.LittleEndian.PutUint16(b[2:], uint16(c.X))
	binary.LittleEndian.PutUint16(b[4:], uint16(c.Y))
	err := common.WriteStringRange(c.Name, b[6:], 1, 20)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c Chest) MarshalBinary() ([]byte, error) {
	return c.Append(make([]byte, 0, c.Len()))
}

func (c *Chest) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.MinLength() {
		return common.ErrInvalidLength
	}

	c.Index = int16(binary.LittleEndian.Uint16(b[0:]))
	c.X = int16(binary.LittleEndian.Uint16(b[2:]))
	c.Y = int16(binary.LittleEndian.Uint16(b[4:]))
	c.Name, err = common.ReadStringRange(b[6:], 1, 20)
	if err != nil {
		return
	}

	return nil
}
