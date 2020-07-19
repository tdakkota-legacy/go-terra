package messages

// Server <-> Client (Sync)
type PlayerHP struct {
	PlayerID byte
	HP       int16
	MaxHP    int16
}
