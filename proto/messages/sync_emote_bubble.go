package messages

// Server -> Client
type SyncEmoteBubble struct {
	EmoteID       int32
	AnchorType    byte
	PlayerID      uint16
	EmoteLifeTime uint16
	Emote         byte
	EmoteMetaData int16
}
