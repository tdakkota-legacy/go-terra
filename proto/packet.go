package proto

//procm:use=derive_binary
type Packet struct {
	Length  uint16
	Tag     Tag
	Message []byte `length:"$m.Length - 2"`
}
