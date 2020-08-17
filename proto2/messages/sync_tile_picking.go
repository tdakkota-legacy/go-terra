package messages

// Client <-> Server
//procm:use=derive_binary
type SyncTilePicking struct {
	PlayerID   byte
	X          int16
	Y          int16
	PickDamage byte
}
