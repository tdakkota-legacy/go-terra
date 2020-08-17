package messages

import "github.com/tdakkota/go-terra/proto2/structs"

// Server <-> Client (Sync)
//procm:use=derive_binary
type SendTileSquare struct {
	Size           uint16
	TileChangeType byte
	TileX          int16
	TileY          int16
	Tiles          []structs.Tile
}
