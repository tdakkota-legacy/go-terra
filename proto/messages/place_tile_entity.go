package messages

// Client -> Server
type PlaceTileEntity struct {
	X              int16
	Y              int16
	TileEntityType byte
}
