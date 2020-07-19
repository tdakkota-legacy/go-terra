package messages

// Server <-> Client (Sync)
type SetActiveNPC struct {
	PlayerID      byte
	NPCTalkTarget int16
}
