package common

import (
	"encoding/binary"
	"math"
)

func WriteFloat32(f float32, buf []byte) {
	binary.LittleEndian.PutUint32(buf, math.Float32bits(f))
}

func ReadFloat32(buf []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(buf))
}
