package messages

// Client -> Server
//procm:use=derive_binary
type PlaceTileEntity struct {
	X              int16
	Y              int16
	TileEntityType byte
}
