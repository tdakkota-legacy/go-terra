package messages

// Client -> Server
type GemLockToggle struct {
	X  int16
	Y  int16
	On bool
}
