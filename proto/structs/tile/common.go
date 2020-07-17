package tile

import (
	"encoding/binary"
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type Dye struct {
	DyeType   int16
	DyePrefix byte
	DyeStack  int16
}

func (d Dye) minLength() int {
	return 2 + 1 + 2
}

func (d Dye) MarshalBinary() ([]byte, error) {
	b := make([]byte, d.minLength())

	binary.LittleEndian.PutUint16(b[0:], uint16(d.DyeType))
	b[2] = byte(d.DyePrefix)
	binary.LittleEndian.PutUint16(b[3:], uint16(d.DyeStack))
	return b, nil
}

func (d *Dye) UnmarshalBinary(b []byte) (err error) {
	if len(b) < d.minLength() {
		return common2.ErrInvalidLength
	}

	d.DyeType = int16(binary.LittleEndian.Uint16(b[0:]))
	d.DyePrefix = byte(b[2])
	d.DyeStack = int16(binary.LittleEndian.Uint16(b[3:]))
	return nil
}

type Item struct {
	ItemType   int16
	ItemPrefix byte
	ItemStack  int16
}

func (i Item) minLength() int {
	return 2 + 1 + 2
}

func (i Item) MarshalBinary() ([]byte, error) {
	b := make([]byte, i.minLength())

	binary.LittleEndian.PutUint16(b[0:], uint16(i.ItemType))
	b[2] = byte(i.ItemPrefix)
	binary.LittleEndian.PutUint16(b[3:], uint16(i.ItemStack))

	return b, nil
}

func (i *Item) UnmarshalBinary(b []byte) (err error) {
	if len(b) < i.minLength() {
		return common2.ErrInvalidLength
	}

	i.ItemType = int16(binary.LittleEndian.Uint16(b[0:]))
	i.ItemPrefix = byte(b[2])
	i.ItemStack = int16(binary.LittleEndian.Uint16(b[3:]))

	return nil
}
