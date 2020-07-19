package messages

// Server <-> Client (Sync)
type SetPlayerStealth struct {
	Player  byte
	Stealth float32
}
