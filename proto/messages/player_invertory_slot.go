package messages

// Server <-> Client (Sync)
type PlayerInventorySlot struct {
	PlayerID  byte
	SlotID    int16
	Stack     int16
	Prefix    byte
	ItemNetID int16
}
