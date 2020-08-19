package messages

// Server -> Client
//procm:use=derive_binary
type AnglerQuest struct {
	Quest     byte
	Completed bool
}
