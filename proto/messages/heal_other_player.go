package messages

// Server <-> Client (Sync)
type HealOtherPlayer struct {
	PlayerID   byte
	HealAmount int16
}
