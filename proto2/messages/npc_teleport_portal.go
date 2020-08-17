package messages

// Server <-> Client
//procm:use=derive_binary
type NPCTeleportPortal struct {
	NPCID            uint16
	PortalColorIndex int16
	NewPositionX     float32
	NewPositionY     float32
	VelocityX        float32
	VelocityY        float32
}
