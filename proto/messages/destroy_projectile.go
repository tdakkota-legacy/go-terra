package messages

// Server <-> Client (Sync)
type DestroyProjectile struct {
	ProjectileID int16
	Owner        byte
}
