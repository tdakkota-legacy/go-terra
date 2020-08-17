package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type HealOtherPlayer struct {
	PlayerID   byte
	HealAmount int16
}
