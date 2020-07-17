package structs

import (
	"encoding/binary"
	common2 "github.com/tdakkota/go-terra/proto/structs/common"
)

type TileFlags1 byte

const (
	Active TileFlags1 = 1 << iota
	Lighted
	HasWall
	HasLiquid
	Wire1
	HalfBrick
	Actuator
	Inactive
)

type TileFlags2 byte

const (
	Wire2 TileFlags2 = 1 << iota
	Wire3
	HasColor
	HasWallColor
	Slope1
	Slope2
	Slope3
	Wire4
)

type Tile struct {
	Flags1     TileFlags1
	Flags2     TileFlags2
	Color      byte
	WallColor  byte
	Type       uint16
	FrameX     int16
	FrameY     int16
	Wall       uint16
	Liquid     byte
	LiquidType LiquidType
}

func (c Tile) minLength() int {
	const flags = 2 * 1                                         // Flags1,Flags2
	const colors = 2 * 1                                        // Color,WallColor
	const _type = 2                                             // Type
	const frameCoords = 2 * 2                                   // FrameX,FrameY
	const wall = 2                                              // Wall
	const liquid = 2 * 1                                        // Liquid,LiquidType
	return flags + colors + _type + frameCoords + wall + liquid // Total
}

func (c Tile) MarshalBinary() ([]byte, error) {
	b := make([]byte, c.minLength())

	b[0], b[1] = byte(c.Flags1), byte(c.Flags2)
	b[2], b[3] = c.Color, c.WallColor
	binary.LittleEndian.PutUint16(b[4:], c.Type)
	binary.LittleEndian.PutUint16(b[6:], uint16(c.FrameX))
	binary.LittleEndian.PutUint16(b[8:], uint16(c.FrameY))
	binary.LittleEndian.PutUint16(b[10:], c.Wall)
	b[12], b[13] = c.Liquid, byte(c.LiquidType)

	return b, nil
}

func (c *Tile) UnmarshalBinary(b []byte) (err error) {
	if len(b) < c.minLength() {
		return common2.ErrInvalidLength
	}

	c.Flags1, c.Flags2 = TileFlags1(b[0]), TileFlags2(b[1])
	c.Color, c.WallColor = b[2], b[3]
	c.Type = binary.LittleEndian.Uint16(b[4:])
	c.FrameX = int16(binary.LittleEndian.Uint16(b[6:]))
	c.FrameY = int16(binary.LittleEndian.Uint16(b[8:]))
	c.Wall = binary.LittleEndian.Uint16(b[10:])
	c.Liquid, c.LiquidType = b[12], LiquidType(b[13])

	return nil
}
