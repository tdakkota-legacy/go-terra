package messages

// Client -> Server
//procm:use=derive_binary
type KillPortal struct {
	ProjectileOwner uint16
	ProjectileAI    byte
}
