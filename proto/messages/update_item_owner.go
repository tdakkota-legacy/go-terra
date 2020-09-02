package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type UpdateItemOwner struct {
	ItemID   int16
	PlayerID byte
}
