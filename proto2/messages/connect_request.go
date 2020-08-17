package messages

// Client -> Server
//procm:use=derive_binary
type ConnectRequest struct {
	Version string
}
