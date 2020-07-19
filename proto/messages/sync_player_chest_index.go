package messages

// Server -> Client
type SyncPlayerChestIndex struct {
	Player byte
	Chest  int16
}
