package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type GrowFX struct {
	EffectFlags byte
	X           int32
	Y           int32
	Data        byte
	TreeGore    int16
}
