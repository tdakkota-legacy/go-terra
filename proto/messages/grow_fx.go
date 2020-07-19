package messages

// Server <-> Client (Sync)
type GrowFX struct {
	EffectFlags byte
	X           int32
	Y           int32
	Data        byte
	TreeGore    int16
}
