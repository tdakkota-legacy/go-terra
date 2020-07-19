package messages

// Server <-> Client (Sync)
type PlayerItemAnimation struct {
	PlayerID      byte
	ItemRotation  float32
	ItemAnimation int16
}
