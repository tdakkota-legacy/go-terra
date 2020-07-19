package messages

// Server -> Client
type CreateTemporaryAnimation struct {
	AnimationType int16
	TileType      uint16
	X             int16
	Y             int16
}
