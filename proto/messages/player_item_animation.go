package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerItemAnimation struct {
	PlayerID      byte
	ItemRotation  float32
	ItemAnimation int16
}
