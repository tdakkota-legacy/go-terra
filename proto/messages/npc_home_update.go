package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type NPCHomeUpdate struct {
	NPCID     int16
	HomeTileX int16
	HomeTileY int16
	Homeless  byte
}
