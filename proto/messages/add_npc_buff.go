package messages

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

// Server <-> Client (Sync)
type AddNPCBuff struct {
	NPCID int16
	Buff  uint16
	Time  int16
}

func (a AddNPCBuff) Len() int {
	return a.MinLength()
}

func (a AddNPCBuff) MinLength() int {
	return 0 + 2 + 2 + 2
}

func (a AddNPCBuff) MarshalBinary() (b []byte, err error) {
	return a.Append(make([]byte, 0, a.Len()))
}

func (a AddNPCBuff) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, a.Len())

	binary.LittleEndian.PutUint16(b[0:], uint16(a.NPCID))
	binary.LittleEndian.PutUint16(b[2:], uint16(a.Buff))
	binary.LittleEndian.PutUint16(b[4:], uint16(a.Time))
	return buf, nil
}

func (a *AddNPCBuff) UnmarshalBinary(b []byte) (err error) {
	if len(b) < a.MinLength() {
		return common.ErrInvalidLength
	}

	a.NPCID = int16(binary.LittleEndian.Uint16(b[0:]))
	a.Buff = uint16(binary.LittleEndian.Uint16(b[2:]))
	a.Time = int16(binary.LittleEndian.Uint16(b[4:]))
	return nil
}
