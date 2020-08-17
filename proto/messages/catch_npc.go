package messages

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

// Client -> Server
type CatchNPC struct {
	NPCID    int16
	PlayerID byte
}

func (c CatchNPC) Len() int {
	return c.MinLength()
}

func (c CatchNPC) MinLength() int {
	return 0 + 2 + 1
}

func (c CatchNPC) MarshalBinary() (b []byte, err error) {
	return c.Append(make([]byte, 0, c.Len()))
}

func (c CatchNPC) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, c.Len())

	binary.LittleEndian.PutUint16(b[0:], uint16(c.NPCID))
	b[2] = byte(c.PlayerID)
	return buf, nil
}

func (c *CatchNPC) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.MinLength() {
		return common.ErrInvalidLength
	}

	c.NPCID = int16(binary.LittleEndian.Uint16(b[0:]))
	c.PlayerID = byte(b[2])
	return nil
}
