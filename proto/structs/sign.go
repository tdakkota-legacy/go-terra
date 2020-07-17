package structs

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

type Sign struct {
	Index    int16
	X        int16
	Y        int16
	SignText string
}

func (s Sign) Len() int {
	return s.MinLength() + len(s.SignText)
}

func (s Sign) MinLength() int {
	return 3*2 + 1 // Index(2) + X(2) + Y(2) + Len(1)
}

func (s Sign) Append(buf []byte) ([]byte, error) {
	var b []byte
	buf, b = common.Slice(buf, s.Len())

	binary.LittleEndian.PutUint16(b[0:], uint16(s.Index))
	binary.LittleEndian.PutUint16(b[2:], uint16(s.X))
	binary.LittleEndian.PutUint16(b[4:], uint16(s.Y))

	err := common.WriteStringRange(s.SignText, b[6:], 1, common.MaxStringLength)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s Sign) MarshalBinary() ([]byte, error) {
	return s.Append(make([]byte, 0, s.Len()))
}

func (s *Sign) UnmarshalBinary(b []byte) (err error) {
	if len(b) < s.MinLength() {
		return common.ErrInvalidLength
	}

	s.Index = int16(binary.LittleEndian.Uint16(b[0:]))
	s.X = int16(binary.LittleEndian.Uint16(b[2:]))
	s.Y = int16(binary.LittleEndian.Uint16(b[4:]))
	s.SignText, err = common.ReadStringRange(b[6:], 1, common.MaxStringLength)
	if err != nil {
		return
	}

	return nil
}
