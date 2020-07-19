package messages

// Client <-> Server
type RequestTileEntityInteraction struct {
	TileEntityID int32
	PlayerID     byte
}
