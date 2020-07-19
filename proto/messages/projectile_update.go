package messages

// Server <-> Client (Sync)
type ProjectileUpdate struct {
	ProjectileID   int16
	PositionX      float32
	PositionY      float32
	VelocityX      float32
	VelocityY      float32
	Owner          byte
	Type           int16
	ProjFlags      byte
	AI0            float32
	AI1            float32
	Damage         int16
	Knockback      float32
	OriginalDamage int16
	ProjUUID       int16
}
