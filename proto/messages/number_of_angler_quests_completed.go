package messages

// Client -> Server
type NumberOfAnglerQuestsCompleted struct {
	PlayerID              byte
	AnglerQuestsCompleted int32
	GolferScore           int32
}
