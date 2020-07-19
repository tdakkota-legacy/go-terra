package messages

// Client <-> Server
type SyncTilePicking struct {
	PlayerID   byte
	X          int16
	Y          int16
	PickDamage byte
}
