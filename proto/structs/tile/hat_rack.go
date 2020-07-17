package tile

import (
	"github.com/tdakkota/go-terra/proto/common"
)

type HatRack struct {
	HatFlags byte
	Item
	Dye
}

func (h HatRack) Len() int {
	return 1 + h.Item.Len() + h.Dye.Len()
}

func (h HatRack) MinLength() int {
	return 0 + 1 + h.Item.MinLength() + h.Dye.MinLength()
}

func (h HatRack) MarshalBinary() (b []byte, err error) {
	return h.Append(make([]byte, 0, h.Len()))
}

func (h HatRack) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, h.Len())
	b = b[:1]

	b[0] = byte(h.HatFlags)
	b, err = h.Item.Append(b)
	if err != nil {
		return nil, err
	}

	b, err = h.Dye.Append(b)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (h *HatRack) UnmarshalBinary(b []byte) (err error) {
	if len(b) < h.MinLength() {
		return common.ErrInvalidLength
	}

	h.HatFlags = byte(b[0])
	err = h.Item.UnmarshalBinary(b[1:])
	if err != nil {
		return
	}

	err = h.Dye.UnmarshalBinary(b[6:])
	if err != nil {
		return
	}

	return nil
}
