package messages

// Server -> Client
//procm:use=derive_binary
type ReportInvasionProgress struct {
	Progress    int32
	MaxProgress int32
	Icon        int8
	Wave        int8
}
