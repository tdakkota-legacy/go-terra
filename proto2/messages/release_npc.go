package messages

// Client -> Server
//procm:use=derive_binary
type ReleaseNPC struct {
	X       int32
	Y       int32
	NPCType int16
	Style   byte
}
