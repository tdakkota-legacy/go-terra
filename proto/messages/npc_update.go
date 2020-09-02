package messages

type NPCUpdateNpcFlags1 byte

const (
	NPCUpdateNpcFlags1Direction       NPCUpdateNpcFlags1 = 1
	NPCUpdateNpcFlags1DirectionY                         = 2
	NPCUpdateNpcFlags1AI0                                = 4
	NPCUpdateNpcFlags1AI1                                = 8
	NPCUpdateNpcFlags1AI2                                = 16
	NPCUpdateNpcFlags1AI3                                = 32
	NPCUpdateNpcFlags1SpriteDirection                    = 64
	NPCUpdateNpcFlags1LifeMax                            = 128
)

type NPCUpdateNpcFlags2 byte

const (
	NPCUpdateNpcFlags2StatsScaled        NPCUpdateNpcFlags2 = 1
	NPCUpdateNpcFlags2SpawnedFromStatue                     = 2
	NPCUpdateNpcFlags2StrengthMultiplier                    = 4
)

// Server -> Client
//procm:use=derive_binary
type NPCUpdate struct {
	NPCID                                       int16
	PositionX                                   float32
	PositionY                                   float32
	VelocityX                                   float32
	VelocityY                                   float32
	Target                                      uint16
	NpcFlags1                                   NPCUpdateNpcFlags1
	NpcFlags2                                   NPCUpdateNpcFlags2
	AI0                                         float32 `if:"$m.NpcFlags1 & NPCUpdateNpcFlags1AI0 != 0"`
	AI1                                         float32 `if:"$m.NpcFlags1 & NPCUpdateNpcFlags1AI1 != 0"`
	AI2                                         float32 `if:"$m.NpcFlags1 & NPCUpdateNpcFlags1AI2 != 0"`
	AI3                                         float32 `if:"$m.NpcFlags1 & NPCUpdateNpcFlags1AI3 != 0"`
	NPCNetID                                    int16
	PlayerCountForMultiplayerDifficultyOverride byte    `if:"$m.NpcFlags2 & NPCUpdateNpcFlags2StatsScaled != 0"`
	StrengthMultiplier                          float32 `if:"$m.NpcFlags2 & NPCUpdateNpcFlags2StrengthMultiplier != 0"`
	LifeBytes                                   byte    `if:"$m.NpcFlags1 & NPCUpdateNpcFlags1LifeMax != 0"`
	Life                                        int32   `if:"$m.NpcFlags1 & NPCUpdateNpcFlags1LifeMax != 0"`
	ReleaseOwner                                byte
}
