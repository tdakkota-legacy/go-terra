package structs

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
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

func (t Tile) Len() int {
	return t.MinLength()
}

func (t Tile) MinLength() int {
	const flags = 2 * 1                                         // Flags1,Flags2
	const colors = 2 * 1                                        // Color,WallColor
	const _type = 2                                             // Type
	const frameCoords = 2 * 2                                   // FrameX,FrameY
	const wall = 2                                              // Wall
	const liquid = 2 * 1                                        // Liquid,LiquidType
	return flags + colors + _type + frameCoords + wall + liquid // Total
}

func (t Tile) Append(buf []byte) ([]byte, error) {
	var b []byte
	buf, b = common.Slice(buf, t.Len())

	b[0], b[1] = byte(t.Flags1), byte(t.Flags2)
	b[2], b[3] = t.Color, t.WallColor
	binary.LittleEndian.PutUint16(b[4:], t.Type)
	binary.LittleEndian.PutUint16(b[6:], uint16(t.FrameX))
	binary.LittleEndian.PutUint16(b[8:], uint16(t.FrameY))
	binary.LittleEndian.PutUint16(b[10:], t.Wall)
	b[12], b[13] = t.Liquid, byte(t.LiquidType)

	return b, nil
}

func (t Tile) MarshalBinary() ([]byte, error) {
	return t.Append(make([]byte, 0, t.Len()))
}

func (t *Tile) UnmarshalBinary(b []byte) (err error) {
	if len(b) < t.MinLength() {
		return common.ErrInvalidLength
	}

	t.Flags1, t.Flags2 = TileFlags1(b[0]), TileFlags2(b[1])
	t.Color, t.WallColor = b[2], b[3]
	t.Type = binary.LittleEndian.Uint16(b[4:])
	t.FrameX = int16(binary.LittleEndian.Uint16(b[6:]))
	t.FrameY = int16(binary.LittleEndian.Uint16(b[8:]))
	t.Wall = binary.LittleEndian.Uint16(b[10:])
	t.Liquid, t.LiquidType = b[12], LiquidType(b[13])

	return nil
}
