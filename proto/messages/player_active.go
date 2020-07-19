package messages

// Server -> Client
type PlayerActive struct {
	PlayerID byte
	Active   byte
}
