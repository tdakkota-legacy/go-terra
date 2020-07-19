package messages

// Server -> Client
type NPCUpdate struct {
	NPCID                                       int16
	PositionX                                   float32
	PositionY                                   float32
	VelocityX                                   float32
	VelocityY                                   float32
	Target                                      uint16
	NpcFlags1                                   byte
	NpcFlags2                                   byte
	AI                                          []float32
	NPCNetID                                    int16
	playerCountForMultiplayerDifficultyOverride byte
	StrengthMultiplier                          float32
	LifeBytes                                   byte
	Life                                        int32
	ReleaseOwner                                byte
}
