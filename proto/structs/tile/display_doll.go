package tile

import (
	"github.com/tdakkota/go-terra/proto/common"
)

type DisplayDoll struct {
	DollFlags1 byte
	DollFlags2 byte
	Item
	Dye
}

func (d DisplayDoll) Len() int {
	return 2 + d.Item.Len() + d.Dye.Len()
}

func (d DisplayDoll) MinLength() int {
	return 2 + d.Item.MinLength() + d.Dye.MinLength()
}

func (d DisplayDoll) Append(buf []byte) (bout []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, d.Len())
	b = b[:2]

	b[0] = byte(d.DollFlags1)
	b[1] = byte(d.DollFlags2)
	b, err = d.Item.Append(b)
	if err != nil {
		return nil, err
	}

	b, err = d.Dye.Append(b)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (d DisplayDoll) MarshalBinary() (b []byte, err error) {
	return d.Append(make([]byte, 0, d.Len()))
}

func (d *DisplayDoll) UnmarshalBinary(b []byte) (err error) {
	if len(b) < d.MinLength() {
		return common.ErrInvalidLength
	}

	d.DollFlags1 = byte(b[0])
	d.DollFlags2 = byte(b[1])
	err = d.Item.UnmarshalBinary(b[2:])
	if err != nil {
		return
	}

	err = d.Dye.UnmarshalBinary(b[7:])
	if err != nil {
		return
	}

	return nil
}
