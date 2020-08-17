package messages

// Client -> Server
//procm:use=derive_binary
type ForceItemIntoNearestChest struct {
	InventorySlot byte
}
