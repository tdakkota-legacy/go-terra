package messages

// Server <-> Client (Sync)
type Unlock struct {
	Type byte
	X    int16
	Y    int16
}
