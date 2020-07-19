package messages

// Server <-> Client (Sync)
type UpdateNPCName struct {
	NPCID                 int16
	Name                  string
	TownNpcVariationIndex int32
}
