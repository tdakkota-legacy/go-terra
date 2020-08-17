package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerMana struct {
	PlayerID byte
	Mana     int16
	MaxMana  int16
}
