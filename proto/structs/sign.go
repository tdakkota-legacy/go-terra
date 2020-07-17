package structs

import (
	"encoding/binary"
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type Sign struct {
	Index    int16
	X        int16
	Y        int16
	SignText string
}

func (c Sign) minLength() int {
	return 3*2 + 1 // Index(2) + X(2) + Y(2) + Len(1)
}

func (c Sign) MarshalBinary() ([]byte, error) {
	b := make([]byte, c.minLength()+len(c.SignText))

	binary.LittleEndian.PutUint16(b[0:], uint16(c.Index))
	binary.LittleEndian.PutUint16(b[2:], uint16(c.X))
	binary.LittleEndian.PutUint16(b[4:], uint16(c.Y))
	err := common2.WriteStringRange(c.SignText, b[6:], 1, common2.MaxStringLength)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Sign) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.minLength() {
		return common2.ErrInvalidLength
	}

	c.Index = int16(binary.LittleEndian.Uint16(b[0:]))
	c.X = int16(binary.LittleEndian.Uint16(b[2:]))
	c.Y = int16(binary.LittleEndian.Uint16(b[4:]))
	c.SignText, err = common2.ReadStringRange(b[6:], 1, common2.MaxStringLength)
	if err != nil {
		return
	}

	return nil
}
