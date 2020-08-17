package messages

// Server -> Client
//procm:use=derive_binary
type CreateTemporaryAnimation struct {
	AnimationType int16
	TileType      uint16
	X             int16
	Y             int16
}
