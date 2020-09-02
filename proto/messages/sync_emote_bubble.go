package messages

// Server -> Client
//procm:use=derive_binary
type SyncEmoteBubble struct {
	EmoteID       int32
	AnchorType    byte
	PlayerID      uint16 `if:"$m.AnchorType != 255"`
	EmoteLifeTime uint16 `if:"$m.AnchorType != 255"`
	Emote         byte   `if:"$m.AnchorType != 255"`
	EmoteMetaData int16  `if:"$m.AnchorType != 255 && $m.Emote < 0"`
}
