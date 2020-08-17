package messages

// Server -> Client
//procm:use=derive_binary
type NotifyPlayerOfEvent struct {
	EventID int16
}
