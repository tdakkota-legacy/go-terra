package messages

// Client -> Server
//procm:use=derive_binary
type GemLockToggle struct {
	X  int16
	Y  int16
	On bool
}
