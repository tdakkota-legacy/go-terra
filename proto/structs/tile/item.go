package tile

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

type Item struct {
	ItemType   int16
	ItemPrefix byte
	ItemStack  int16
}

func (i Item) Len() int {
	return i.MinLength()
}

func (i Item) MinLength() int {
	return 0 + 2 + 1 + 2
}

func (i Item) MarshalBinary() (b []byte, err error) {
	return i.Append(make([]byte, 0, i.Len()))
}

func (i Item) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, i.Len())

	binary.LittleEndian.PutUint16(b[0:], uint16(i.ItemType))
	b[2] = byte(i.ItemPrefix)
	binary.LittleEndian.PutUint16(b[3:], uint16(i.ItemStack))
	return buf, nil
}

func (i *Item) UnmarshalBinary(b []byte) (err error) {
	if len(b) < i.MinLength() {
		return common.ErrInvalidLength
	}

	i.ItemType = int16(binary.LittleEndian.Uint16(b[0:]))
	i.ItemPrefix = byte(b[2])
	i.ItemStack = int16(binary.LittleEndian.Uint16(b[3:]))
	return nil
}
