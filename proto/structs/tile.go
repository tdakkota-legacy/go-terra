package structs

//procm:use=derive_binary
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

//procm:use=derive_binary
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

//procm:use=derive_binary
type Tile struct {
	Flags1     TileFlags1
	Flags2     TileFlags2
	Color      byte       `if:"$m.Flags2 & HasColor != 0"`
	WallColor  byte       `if:"$m.Flags2 & HasWallColor != 0"`
	Type       uint16     `if:"$m.Flags1 & Active != 0"`
	FrameX     int16      `if:"$m.Flags1 & Active != 0"`
	FrameY     int16      `if:"$m.Flags1 & Active != 0"`
	Wall       uint16     `if:"$m.Flags1 & HasWall != 0"`
	Liquid     byte       `if:"$m.Flags1 & HasLiquid != 0"`
	LiquidType LiquidType `if:"$m.Flags1 & HasLiquid != 0"`
}
