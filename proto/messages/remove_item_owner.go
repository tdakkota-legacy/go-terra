package messages

// Server -> Client
//procm:use=derive_binary
type RemoveItemOwner struct {
	ItemIndex int16
}
