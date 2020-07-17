package tile

import (
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type HatRack struct {
	HatFlags byte
	Item
	Dye
}

func (h HatRack) minLength() int {
	return 1 + h.Item.minLength() + h.Dye.minLength()
}

func (h HatRack) MarshalBinary() ([]byte, error) {
	b := make([]byte, h.minLength())

	b[0] = byte(h.HatFlags)
	dataItem, err := h.Item.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(dataItem, b[1:])

	dataDye, err := h.Dye.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(dataDye, b[6:])

	return b, nil
}

func (h *HatRack) UnmarshalBinary(b []byte) (err error) {
	if len(b) < h.minLength() {
		return common2.ErrInvalidLength
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
