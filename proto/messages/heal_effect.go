package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type HealEffect struct {
	PlayerID   byte
	HealAmount int16
}
