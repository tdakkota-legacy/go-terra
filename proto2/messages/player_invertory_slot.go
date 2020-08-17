package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerInventorySlot struct {
	PlayerID  byte
	SlotID    int16
	Stack     int16
	Prefix    byte
	ItemNetID int16
}
