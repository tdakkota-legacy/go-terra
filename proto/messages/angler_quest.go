package messages

import "github.com/tdakkota/go-terra/proto/common"

// Server -> Client
type AnglerQuest struct {
	Quest     byte
	Completed bool
}

func (a AnglerQuest) Len() int {
	return a.MinLength()
}

func (a AnglerQuest) MinLength() int {
	return 0 + 1 + 1
}

func (a AnglerQuest) MarshalBinary() (b []byte, err error) {
	return a.Append(make([]byte, 0, a.Len()))
}

func (a AnglerQuest) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, a.Len())

	b[0] = byte(a.Quest)
	common.WriteBool(a.Completed, b[1:])
	return buf, nil
}

func (a *AnglerQuest) UnmarshalBinary(b []byte) (err error) {
	if len(b) < a.MinLength() {
		return common.ErrInvalidLength
	}

	a.Quest = byte(b[0])
	a.Completed = common.ReadBool(b[1:])
	return nil
}
