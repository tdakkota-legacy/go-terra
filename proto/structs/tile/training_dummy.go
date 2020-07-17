package tile

import (
	"encoding/binary"
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type TrainingDummy struct {
	NPCIndex int16
}

func (h TrainingDummy) minLength() int {
	return 2
}

func (h TrainingDummy) MarshalBinary() ([]byte, error) {
	b := make([]byte, h.minLength())

	binary.LittleEndian.PutUint16(b[0:], uint16(h.NPCIndex))
	return b, nil
}

func (h *TrainingDummy) UnmarshalBinary(b []byte) (err error) {
	if len(b) < h.minLength() {
		return common2.ErrInvalidLength
	}

	h.NPCIndex = int16(binary.LittleEndian.Uint16(b[0:]))
	return nil
}
