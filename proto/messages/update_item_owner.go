package messages

// Server <-> Client (Sync)
type UpdateItemOwner struct {
	ItemID   int16
	PlayerID byte
}
