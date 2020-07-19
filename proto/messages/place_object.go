package messages

// Server <-> Client
type PlaceObject struct {
	X         int16
	Y         int16
	Type      int16
	Style     int16
	Alternate byte
	Random    int8
	Direction bool
}
