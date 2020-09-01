package messages

import "github.com/tdakkota/go-terra/proto/structs"

// Server <-> Client (Sync)
//procm:use=derive_binary
type SendTileSquare struct {
	Size           uint16
	TileChangeType byte `if:"$m.Size & 0x8000 != 0"`
	TileX          int16
	TileY          int16
	Tiles          []structs.Tile `length:"$m.Size"`
}
