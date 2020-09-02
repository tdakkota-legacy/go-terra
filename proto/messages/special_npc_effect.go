package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type SpecialNPCEffect struct {
	PlayerID byte
	Type     byte
}
