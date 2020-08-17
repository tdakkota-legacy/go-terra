package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type SetActiveNPC struct {
	PlayerID      byte
	NPCTalkTarget int16
}
