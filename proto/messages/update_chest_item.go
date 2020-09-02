package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type UpdateChestItem struct {
	ChestID   int16
	ItemSlot  byte
	Stack     int16
	Prefix    byte
	ItemNetID int16
}
