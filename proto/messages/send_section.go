package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Server -> Client
//procm:use=derive_binary
type SendSection struct {
	Compressed      bool
	XStart          int32
	YStart          int32
	Width           int16
	Height          int16
	Tiles           []structs.Tile
	ChestCount      int16
	Chests          []structs.Chest
	SignCount       int16
	Signs           []structs.Sign
	TileEntityCount int16
	TileEntities    []structs.TileEntity
}
