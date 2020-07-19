package messages

// Server <-> Client (Sync)
type TogglePVP struct {
	PlayerID   byte
	PVPEnabled bool
}
