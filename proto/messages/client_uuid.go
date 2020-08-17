package messages

import "github.com/tdakkota/go-terra/proto/common"

// Client -> Server
type ClientUUID struct {
	UUID string
}

func (c ClientUUID) Len() int {
	return c.MinLength() + len(c.UUID)
}

func (c ClientUUID) MinLength() int {
	return 0 + 1
}

func (c ClientUUID) MarshalBinary() (b []byte, err error) {
	return c.Append(make([]byte, 0, c.Len()))
}

func (c ClientUUID) Append(buf []byte) (_ []byte, err error) {
	var b []byte
	buf, b = common.Slice(buf, c.Len())

	err = common.WriteString(c.UUID, b[0:])
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *ClientUUID) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.MinLength() {
		return common.ErrInvalidLength
	}

	c.UUID, err = common.ReadString(b[0:])
	if err != nil {
		return
	}

	return nil
}
