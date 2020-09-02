package messages

// Updates sign if sent from client otherwise displays sign to client.
//procm:use=derive_binary
type UpdateSign struct {
	SignID    int16
	X         int16
	Y         int16
	Text      string
	PlayerID  byte
	SignFlags byte
}
