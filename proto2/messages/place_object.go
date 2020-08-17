package messages

// Server <-> Client
//procm:use=derive_binary
type PlaceObject struct {
	X         int16
	Y         int16
	Type      int16
	Style     int16
	Alternate byte
	Random    int8
	Direction bool
}
