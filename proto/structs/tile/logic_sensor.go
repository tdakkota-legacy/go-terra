package tile

import (
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type LogicSensor struct {
	LogicCheckType byte
	On             bool
}

func (l LogicSensor) minLength() int {
	return 2
}

func (l LogicSensor) MarshalBinary() (b []byte, err error) {
	b = make([]byte, l.minLength())

	b[0] = byte(l.LogicCheckType)
	common2.WriteBool(l.On, b[1:])

	return b, nil
}

func (l *LogicSensor) UnmarshalBinary(b []byte) (err error) {
	if len(b) < l.minLength() {
		return common2.ErrInvalidLength
	}

	l.LogicCheckType = byte(b[0])
	l.On = common2.ReadBool(b[1:])

	return nil
}
