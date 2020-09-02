package messages

// Client -> Server
//procm:use=derive_binary
type NumberOfAnglerQuestsCompleted struct {
	PlayerID              byte
	AnglerQuestsCompleted int32
	GolferScore           int32
}
