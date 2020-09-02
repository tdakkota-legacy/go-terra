package messages

// Server -> Client
//procm:use=derive_binary
type UpdateTileEntity struct {
	TileEntityId   int32
	UpdateTileFlag bool
	TileEntityType byte  `if:"$m.UpdateTileFlag"`
	X              int16 `if:"$m.UpdateTileFlag"`
	Y              int16 `if:"$m.UpdateTileFlag"`
}
