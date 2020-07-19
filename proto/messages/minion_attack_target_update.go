package messages

// Client -> Server
type MinionAttackTargetUpdate struct {
	PlayerID           byte
	MinionAttackTarget int16
}
