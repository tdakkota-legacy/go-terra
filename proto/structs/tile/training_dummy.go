package tile

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

type TrainingDummy struct {
	NPCIndex int16
}

func (t TrainingDummy) Len() int {
	return t.MinLength()
}

func (t TrainingDummy) MinLength() int {
	return 0 + 2
}

func (t TrainingDummy) MarshalBinary() (b []byte, err error) {
	return t.Append(make([]byte, 0, t.Len()))
}

func (t TrainingDummy) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, t.Len())

	binary.LittleEndian.PutUint16(b[0:], uint16(t.NPCIndex))
	return buf, nil
}

func (t *TrainingDummy) UnmarshalBinary(b []byte) (err error) {
	if len(b) < t.MinLength() {
		return common.ErrInvalidLength
	}

	t.NPCIndex = int16(binary.LittleEndian.Uint16(b[0:]))
	return nil
}
