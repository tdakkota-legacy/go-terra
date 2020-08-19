package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type PlayerTeam struct {
	PlayerID byte
	Team     byte
}
