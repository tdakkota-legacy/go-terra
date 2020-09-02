package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerDodge struct {
	PlayerID byte
	Flag     byte
}
