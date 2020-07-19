package messages

// Server -> Client
type NPCShopItem struct {
	Slot     byte
	ItemType int16
	Stack    int16
	Prefix   byte
	Value    int32
	Flags    byte
}
