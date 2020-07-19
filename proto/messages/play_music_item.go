package messages

// Server <-> Client (Sync)
type PlayMusicItem struct {
	PlayerID byte
	Note     float32
}
