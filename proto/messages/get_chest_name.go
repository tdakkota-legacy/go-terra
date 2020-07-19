package messages

// Server <-> Client (Sync)
type GetChestName struct {
	ChestID int16
	ChestX  int16
	ChestY  int16
	Name    string
}
