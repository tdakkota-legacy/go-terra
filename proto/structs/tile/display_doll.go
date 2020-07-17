package tile

import (
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type DisplayDoll struct {
	DollFlags1 byte
	DollFlags2 byte
	Item
	Dye
}

func (d DisplayDoll) minLength() int {
	return 1 + d.Item.minLength() + d.Dye.minLength()
}

func (d DisplayDoll) MarshalBinary() (b []byte, err error) {
	b = make([]byte, d.minLength())

	b[0] = byte(d.DollFlags1)
	b[1] = byte(d.DollFlags2)
	dataItem, err := d.Item.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(dataItem, b[2:])

	dataDye, err := d.Dye.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(dataDye, b[2:])

	return b, nil
}

func (d *DisplayDoll) UnmarshalBinary(b []byte) (err error) {
	if len(b) < d.minLength() {
		return common2.ErrInvalidLength
	}

	d.DollFlags1 = byte(b[0])
	d.DollFlags2 = byte(b[1])
	err = d.Item.UnmarshalBinary(b[2:])
	if err != nil {
		return
	}

	err = d.Dye.UnmarshalBinary(b[2:])
	if err != nil {
		return
	}

	return nil
}
