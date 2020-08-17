package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type UpdateItemDrop2 struct {
	ItemID    int16
	PositionX float32
	PositionY float32
	VelocityX float32
	VelocityY float32
	StackSize int16
	Prefix    byte
	NoDelay   byte
	ItemNetID int16
}
