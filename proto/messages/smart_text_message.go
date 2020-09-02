package messages

import (
	"github.com/tdakkota/go-terra/proto/structs"
)

// Server -> Client
//procm:use=derive_binary
type SmartTextMessage struct {
	MessageColor  structs.Color
	Message       structs.NetworkText
	MessageLength int16
}
