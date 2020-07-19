package messages

// Server <-> Client (Sync)
type HealEffect struct {
	PlayerID   byte
	HealAmount int16
}
