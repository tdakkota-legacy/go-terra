package messages

// Client -> Server
//procm:use=derive_binary
type Emoji struct {
	PlayerIndex byte
	EmoticonID  byte
}
