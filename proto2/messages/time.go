package messages

// Server -> Client
//procm:use=derive_binary
type Time struct {
	DayTime   bool
	TimeValue int32
	SunModY   int16
	MoonModY  int16
}
