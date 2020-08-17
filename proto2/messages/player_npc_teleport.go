package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerNPCTeleport struct {
	Flags     byte
	TargetID  int16
	X         float32
	Y         float32
	Style     byte
	ExtraInfo int32
}
