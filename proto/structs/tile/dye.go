package tile

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

type Dye struct {
	DyeType   int16
	DyePrefix byte
	DyeStack  int16
}

func (d Dye) Len() int {
	return d.MinLength()
}

func (d Dye) MinLength() int {
	return 0 + 2 + 1 + 2
}

func (d Dye) MarshalBinary() (b []byte, err error) {
	return d.Append(make([]byte, 0, d.Len()))
}

func (d Dye) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, d.Len())

	binary.LittleEndian.PutUint16(b[0:], uint16(d.DyeType))
	b[2] = byte(d.DyePrefix)
	binary.LittleEndian.PutUint16(b[3:], uint16(d.DyeStack))
	return buf, nil
}

func (d *Dye) UnmarshalBinary(b []byte) (err error) {
	if len(b) < d.MinLength() {
		return common.ErrInvalidLength
	}

	d.DyeType = int16(binary.LittleEndian.Uint16(b[0:]))
	d.DyePrefix = byte(b[2])
	d.DyeStack = int16(binary.LittleEndian.Uint16(b[3:]))
	return nil
}
