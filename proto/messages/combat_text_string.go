package messages

import (
	"github.com/tdakkota/go-terra/proto/common"
	"github.com/tdakkota/go-terra/proto/structs"
)

// Client <-> Server
type CombatTextString struct {
	X          float32
	Y          float32
	Color      structs.Color
	CombatText structs.NetworkText
}

func (c CombatTextString) Len() int {
	return c.MinLength() + c.Color.Len() + c.CombatText.Len()
}

func (c CombatTextString) MinLength() int {
	return 0 + 4 + 4
}

func (c CombatTextString) MarshalBinary() (b []byte, err error) {
	return c.Append(make([]byte, 0, c.Len()))
}

func (c CombatTextString) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, c.Len())

	common.WriteFloat32(c.X, b[0:])
	common.WriteFloat32(c.Y, b[4:])
	b = b[:8]

	b, err = c.Color.Append(b)
	if err != nil {
		return nil, err
	}

	b, err = c.CombatText.Append(b)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *CombatTextString) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.MinLength() {
		return common.ErrInvalidLength
	}

	c.X = common.ReadFloat32(b[0:])
	c.Y = common.ReadFloat32(b[4:])

	return nil
}
