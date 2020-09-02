package messages

// Client -> Server
//procm:use=derive_binary
type PlaceItemFrame struct {
	X      int16
	Y      int16
	ItemId int16
	Prefix byte
	Stack  int16
}
