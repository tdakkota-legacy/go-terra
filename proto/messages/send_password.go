package messages

// Client -> Server
//procm:use=derive_binary
type SendPassword struct {
	Password string
}
