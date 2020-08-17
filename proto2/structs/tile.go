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
	Color      byte
	WallColor  byte
	Type       uint16
	FrameX     int16
	FrameY     int16
	Wall       uint16
	Liquid     byte
	LiquidType LiquidType
}
