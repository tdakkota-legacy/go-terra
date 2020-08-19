package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayMusicItem struct {
	PlayerID byte
	Note     float32
}
