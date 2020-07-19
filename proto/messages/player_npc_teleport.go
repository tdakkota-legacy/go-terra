package messages

// Server <-> Client (Sync)
type PlayerNPCTeleport struct {
	Flags     byte
	TargetID  int16
	X         float32
	Y         float32
	Style     byte
	ExtraInfo int32
}
