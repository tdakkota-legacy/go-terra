package messages

import "github.com/tdakkota/go-terra/proto/common"

// Client -> Server
type ClientFinishedInventoryChangesOnThisTick struct{}

func (c ClientFinishedInventoryChangesOnThisTick) Len() int {
	return c.MinLength()
}

func (c ClientFinishedInventoryChangesOnThisTick) MinLength() int {
	return 0
}

func (c ClientFinishedInventoryChangesOnThisTick) MarshalBinary() (b []byte, err error) {
	return
}

func (c ClientFinishedInventoryChangesOnThisTick) Append(buf []byte) (_ []byte, err error) {
	return buf, nil
}

func (c *ClientFinishedInventoryChangesOnThisTick) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.MinLength() {
		return common.ErrInvalidLength
	}

	return nil
}
