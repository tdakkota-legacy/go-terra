package messages

// Server -> Client
type PlayLegacySound struct {
	X          float32
	Y          float32
	SoundID    uint16
	SoundFlags byte
}
