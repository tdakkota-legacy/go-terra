package messages

// Server -> Client
//procm:use=derive_binary
type SetCountsAsHostForGameplay struct {
	PlayerID     byte
	CountsAsHost bool
}
