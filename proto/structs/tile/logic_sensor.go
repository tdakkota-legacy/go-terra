package tile

import (
	"github.com/tdakkota/go-terra/proto/common"
)

type LogicSensor struct {
	LogicCheckType byte
	On             bool
}

func (l LogicSensor) Len() int {
	return l.MinLength()
}

func (l LogicSensor) MinLength() int {
	return 0 + 1 + 1
}

func (l LogicSensor) MarshalBinary() (b []byte, err error) {
	return l.Append(make([]byte, 0, l.Len()))
}

func (l LogicSensor) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, l.Len())

	b[0] = byte(l.LogicCheckType)
	common.WriteBool(l.On, b[1:])

	return buf, nil
}

func (l *LogicSensor) UnmarshalBinary(b []byte) (err error) {
	if len(b) < l.MinLength() {
		return common.ErrInvalidLength
	}

	l.LogicCheckType = byte(b[0])
	l.On = common.ReadBool(b[1:])

	return nil
}
