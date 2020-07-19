package messages

// Server -> Client
type SetCountsAsHostForGameplay struct {
	PlayerID     byte
	CountsAsHost bool
}
