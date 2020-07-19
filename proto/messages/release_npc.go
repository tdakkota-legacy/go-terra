package messages

// Client -> Server
type ReleaseNPC struct {
	X       int32
	Y       int32
	NPCType int16
	Style   byte
}
