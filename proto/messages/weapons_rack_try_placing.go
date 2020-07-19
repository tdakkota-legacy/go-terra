package messages

// Client -> Server
type WeaponsRackTryPlacing struct {
	X      int16
	Y      int16
	NetID  int16
	Prefix byte
	Stack  int16
}
