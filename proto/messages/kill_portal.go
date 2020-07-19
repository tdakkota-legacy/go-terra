package messages

// Client -> Server
type KillPortal struct {
	ProjectileOwner uint16
	ProjectileAI    byte
}
