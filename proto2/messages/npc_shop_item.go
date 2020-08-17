package messages

// Server -> Client
//procm:use=derive_binary
type NPCShopItem struct {
	Slot     byte
	ItemType int16
	Stack    int16
	Prefix   byte
	Value    int32
	Flags    byte
}
