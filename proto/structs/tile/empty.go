package tile

type Empty struct{}

func (e Empty) Len() int {
	return 0
}

func (e Empty) MarshalBinary() (data []byte, err error) {
	return
}

func (e Empty) UnmarshalBinary(data []byte) error {
	return nil
}
