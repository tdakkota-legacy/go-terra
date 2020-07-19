package messages

// Server <-> Client (Sync)
type ManaEffect struct {
	PlayerID   byte
	ManaAmount int16
}
