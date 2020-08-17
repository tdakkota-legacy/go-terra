package messages

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
)

// Server <-> Client (Sync)
type AddPlayerBuff struct {
	PlayerID byte
	Buff     uint16
	Time     int32
}

func (a AddPlayerBuff) Len() int {
	return a.MinLength()
}

func (a AddPlayerBuff) MinLength() int {
	return 0 + 1 + 2 + 4
}

func (a AddPlayerBuff) MarshalBinary() (b []byte, err error) {
	return a.Append(make([]byte, 0, a.Len()))
}

func (a AddPlayerBuff) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, a.Len())

	b[0] = byte(a.PlayerID)
	binary.LittleEndian.PutUint16(b[1:], uint16(a.Buff))
	binary.LittleEndian.PutUint32(b[3:], uint32(a.Time))
	return buf, nil
}

func (a *AddPlayerBuff) UnmarshalBinary(b []byte) (err error) {
	if len(b) < a.MinLength() {
		return common.ErrInvalidLength
	}

	a.PlayerID = byte(b[0])
	a.Buff = uint16(binary.LittleEndian.Uint16(b[1:]))
	a.Time = int32(binary.LittleEndian.Uint32(b[3:]))
	return nil
}
