package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type UpdateNPCName struct {
	NPCID                 int16
	Name                  string
	TownNpcVariationIndex int32
}
