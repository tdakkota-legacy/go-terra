package messages

// Server -> Client
//procm:use=derive_binary
type PlayerActive struct {
	PlayerID byte
	Active   byte
}
