package messages

// Server <-> Client (Sync)
type SyncActiveChest struct {
	ChestID    int16
	ChestX     int16
	ChestY     int16
	NameLength byte
	ChestName  string
}
