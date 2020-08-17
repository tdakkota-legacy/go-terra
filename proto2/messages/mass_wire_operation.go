package messages

// Client -> Server
//procm:use=derive_binary
type MassWireOperation struct {
	StartX   int16
	StartY   int16
	EndX     int16
	EndY     int16
	ToolMode byte
}
