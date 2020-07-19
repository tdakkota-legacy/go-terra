package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Server <-> Client (Sync)
type SendTileSquare struct {
	Size           uint16
	TileChangeType byte
	TileX          int16
	TileY          int16
	Tiles          []structs.Tile
}
