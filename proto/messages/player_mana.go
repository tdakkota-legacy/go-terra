package messages

// Server <-> Client (Sync)
type PlayerMana struct {
	PlayerID byte
	Mana     int16
	MaxMana  int16
}
