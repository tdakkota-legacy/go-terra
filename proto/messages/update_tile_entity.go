package messages

// Server -> Client
type UpdateTileEntity struct {
	TileEntityId   int32
	UpdateTileFlag bool
	TileEntityType byte
	X              int16
	Y              int16
}
