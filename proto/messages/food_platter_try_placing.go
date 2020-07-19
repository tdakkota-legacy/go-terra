package messages

// Client -> Server
type FoodPlatterTryPlacing struct {
	X      int16
	Y      int16
	ItemID int16
	Prefix byte
	Stack  int16
}
