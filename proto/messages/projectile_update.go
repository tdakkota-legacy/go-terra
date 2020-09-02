package messages

type ProjectileUpdateProjFlags byte

const (
	ProjectileUpdateProjFlagsAI0            = 1
	ProjectileUpdateProjFlagsAI1            = 2
	ProjectileUpdateProjFlagsDamage         = 16
	ProjectileUpdateProjFlagsKnockback      = 32
	ProjectileUpdateProjFlagsOriginalDamage = 64
	ProjectileUpdateProjFlagsProjUUID       = 128
)

// Server <-> Client (Sync)
//procm:use=derive_binary
type ProjectileUpdate struct {
	ProjectileID   int16
	PositionX      float32
	PositionY      float32
	VelocityX      float32
	VelocityY      float32
	Owner          byte
	Type           int16
	ProjFlags      byte
	AI0            float32 `if:"$m.ProjFlags & ProjectileUpdateProjFlagsAI0 != 0"`
	AI1            float32 `if:"$m.ProjFlags & ProjectileUpdateProjFlagsAI1 != 0"`
	Damage         int16   `if:"$m.ProjFlags & ProjectileUpdateProjFlagsDamage != 0"`
	Knockback      float32 `if:"$m.ProjFlags & ProjectileUpdateProjFlagsKnockback != 0"`
	OriginalDamage int16   `if:"$m.ProjFlags & ProjectileUpdateProjFlagsOriginalDamage != 0"`
	ProjUUID       int16   `if:"$m.ProjFlags & ProjectileUpdateProjFlagsProjUUID != 0"`
}
