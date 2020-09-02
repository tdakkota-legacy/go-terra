package messages

// Server -> Client
//procm:use=derive_binary
type PlayLegacySound struct {
	X          float32
	Y          float32
	SoundID    uint16
	SoundFlags byte
}
