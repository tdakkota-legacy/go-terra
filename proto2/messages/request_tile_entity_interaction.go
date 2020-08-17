package messages

// Client <-> Server
//procm:use=derive_binary
type RequestTileEntityInteraction struct {
	TileEntityID int32
	PlayerID     byte
}
