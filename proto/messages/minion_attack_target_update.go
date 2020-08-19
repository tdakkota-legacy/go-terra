package messages

// Client -> Server
//procm:use=derive_binary
type MinionAttackTargetUpdate struct {
	PlayerID           byte
	MinionAttackTarget int16
}
