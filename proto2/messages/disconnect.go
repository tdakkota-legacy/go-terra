package messages

import (
	"github.com/tdakkota/go-terra/proto2/structs"
)

// Server -> Client
//procm:use=derive_binary
type Disconnect struct {
	Reason structs.NetworkText
}
