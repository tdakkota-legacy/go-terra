package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type ManaEffect struct {
	PlayerID   byte
	ManaAmount int16
}
