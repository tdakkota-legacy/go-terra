package structs

import (
	"encoding/binary"
	"github.com/tdakkota/go-terra/proto/common"
	"github.com/tdakkota/go-terra/proto/structs/tile"
)

type TileEntity struct {
	Type      TileEntityType
	ID        int32
	X         int16
	Y         int16
	ExtraData TileEntityExtraData
}

func (t TileEntity) Len() int {
	return t.MinLength() + t.ExtraData.Len()
}

func (t TileEntity) MinLength() int {
	return 0 + 1 + 4 + 2 + 2
}

func (t TileEntity) MarshalBinary() (b []byte, err error) {
	return t.Append(make([]byte, 0, t.Len()))
}

func (t *TileEntity) setExtraData(typ TileEntityType) {
	switch typ {
	case TrainingDummy:
		t.ExtraData = &tile.TrainingDummy{}
	case ItemFrame:
		t.ExtraData = &tile.ItemFrame{}
	case LogicSensor:
		t.ExtraData = &tile.LogicSensor{}
	case DisplayDoll:
		t.ExtraData = &tile.DisplayDoll{}
	case WeaponsRack:
		t.ExtraData = &tile.WeaponsRack{}
	case HatRack:
		t.ExtraData = &tile.HatRack{}
	case FloodPlatter:
		t.ExtraData = &tile.FoodPlatter{}
	case TeleportationPylon:
		t.ExtraData = tile.Empty{}
	default:
		t.ExtraData = tile.Empty{}
	}
}

func (t TileEntity) Append(buf []byte) (_ []byte, err error) {
	t.setExtraData(t.Type)
	var b []byte
	buf, b = common.Slice(buf, t.Len())

	b[0] = byte(t.Type)
	binary.LittleEndian.PutUint32(b[1:], uint32(t.ID))
	binary.LittleEndian.PutUint16(b[5:], uint16(t.X))
	binary.LittleEndian.PutUint16(b[7:], uint16(t.Y))
	dataExtraData, err := t.ExtraData.MarshalBinary()
	if err != nil {
		return nil, err
	}
	copy(dataExtraData, b[9:])

	return buf, nil
}

func (t *TileEntity) UnmarshalBinary(b []byte) (err error) {
	if len(b) < t.MinLength() {
		return common.ErrInvalidLength
	}

	t.Type = TileEntityType(b[0])
	t.ID = int32(binary.LittleEndian.Uint32(b[1:]))
	t.X = int16(binary.LittleEndian.Uint16(b[5:]))
	t.Y = int16(binary.LittleEndian.Uint16(b[7:]))

	t.setExtraData(t.Type)
	err = t.ExtraData.UnmarshalBinary(b[9:])
	if err != nil {
		return
	}

	return nil
}
