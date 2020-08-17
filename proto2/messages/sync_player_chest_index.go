package messages

// Server -> Client
//procm:use=derive_binary
type SyncPlayerChestIndex struct {
	Player byte
	Chest  int16
}
