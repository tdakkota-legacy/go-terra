package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type HitSwitch struct {
	X int16
	Y int16
}
