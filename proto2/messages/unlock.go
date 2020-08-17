package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type Unlock struct {
	Type byte
	X    int16
	Y    int16
}
