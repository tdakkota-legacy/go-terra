package messages

// Server <-> Client (Sync)
type NPCStrike struct {
	NPCID        int16
	Damage       int16
	Knockback    float32
	HitDirection byte
	Crit         bool
}
