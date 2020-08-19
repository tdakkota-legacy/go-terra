package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type SetPlayerStealth struct {
	Player  byte
	Stealth float32
}
