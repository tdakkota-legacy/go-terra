package messages

// Client -> Server
//procm:use=derive_binary
type RequestEssentialTiles struct {
	X int32
	Y int32
}
