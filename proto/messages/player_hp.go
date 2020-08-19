package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerHP struct {
	PlayerID byte
	HP       int16
	MaxHP    int16
}
