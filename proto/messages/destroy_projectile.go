package messages

// Server <-> Client (Sync)
//procm:use=derive_binary
type DestroyProjectile struct {
	ProjectileID int16
	Owner        byte
}
