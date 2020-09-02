package messages

// Server -> Client
//procm:use=derive_binary
type UpdateNPCBuff struct {
	NPCID   int16
	BuffID1 uint16
	Time1   int16
	BuffID2 uint16
	Time2   int16
	BuffID3 uint16
	Time3   int16
	BuffID4 uint16
	Time4   int16
	BuffID5 uint16
	Time5   int16
}
