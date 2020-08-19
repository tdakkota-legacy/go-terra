package messages

// Server -> Client
//procm:use=derive_binary
type SetUserSlot struct {
	PlayerID byte
}
