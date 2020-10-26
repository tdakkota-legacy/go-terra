// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m WorldInfo) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt32(m.Time)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.DayandMoonInfo)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.MoonPhase)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.MaxTilesX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.MaxTilesY)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.SpawnX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.SpawnY)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.WorldSurface)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.RockLayer)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.WorldID)
		if err != nil {
			return err
		}
	}
	{
		_, err := cur.WriteString(m.WorldName)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.GameMode)
		if err != nil {
			return err
		}
	}
	{
		for _, v := range m.WorldUniqueID {
			{
				err := cur.WriteUint8(v)
				if err != nil {
					return err
				}
			}
		}
	}
	{
		err := cur.WriteUint64(m.WorldGeneratorVersion)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.MoonType)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.TreeBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CorruptionBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.JungleBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.SnowBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.HallowBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CrimsonBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.DesertBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.OceanBackground)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.UnknownBackground1)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.UnknownBackground2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.UnknownBackground3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.UnknownBackground4)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.UnknownBackground5)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.IceBackStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.JungleBackStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.HellBackStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.WindSpeedSet)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CloudNumber)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.Tree1)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.Tree2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.Tree3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.TreeStyle1)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.TreeStyle2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.TreeStyle3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.TreeStyle4)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.CaveBack1)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.CaveBack2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.CaveBack3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CaveBackStyle1)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CaveBackStyle2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CaveBackStyle3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CaveBackStyle4)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Forest1TreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Forest2TreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Forest3TreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Forest4TreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CorruptionTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.JungleTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.SnowTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.HallowTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.CrimsonTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.DesertTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.OceanTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.GlowingMushroomTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.UnderworldTreeTopStyle)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.Rain)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.EventInfo)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.EventInfo2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.EventInfo3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.EventInfo4)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.EventInfo5)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.EventInfo6)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.EventInfo7)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.CopperOreTier)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.IronOreTier)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.SilverOreTier)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.GoldOreTier)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.CobaltOreTier)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.MythrilOreTier)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.AdamantiteOreTier)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt8(m.InvasionType)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint64(m.LobbyID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.SandstormSeverity)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *WorldInfo) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.Time = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.DayandMoonInfo = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.MoonPhase = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.MaxTilesX = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.MaxTilesY = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.SpawnX = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.SpawnY = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.WorldSurface = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.RockLayer = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.WorldID = tmp
	}
	{
		tmp, err := cur.ReadString()
		if err != nil {
			return err
		}
		m.WorldName = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.GameMode = tmp
	}
	{
		n := 16
		for i := 0; i < int(n); i++ {
			{
				tmp, err := cur.ReadUint8()
				if err != nil {
					return err
				}
				m.WorldUniqueID[i] = tmp
			}
		}
	}
	{
		tmp, err := cur.ReadUint64()
		if err != nil {
			return err
		}
		m.WorldGeneratorVersion = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.MoonType = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.TreeBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CorruptionBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.JungleBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.SnowBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.HallowBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CrimsonBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.DesertBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.OceanBackground = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.UnknownBackground1 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.UnknownBackground2 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.UnknownBackground3 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.UnknownBackground4 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.UnknownBackground5 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.IceBackStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.JungleBackStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.HellBackStyle = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.WindSpeedSet = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CloudNumber = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.Tree1 = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.Tree2 = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.Tree3 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.TreeStyle1 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.TreeStyle2 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.TreeStyle3 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.TreeStyle4 = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.CaveBack1 = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.CaveBack2 = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.CaveBack3 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CaveBackStyle1 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CaveBackStyle2 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CaveBackStyle3 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CaveBackStyle4 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Forest1TreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Forest2TreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Forest3TreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Forest4TreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CorruptionTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.JungleTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.SnowTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.HallowTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.CrimsonTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.DesertTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.OceanTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.GlowingMushroomTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.UnderworldTreeTopStyle = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.Rain = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.EventInfo = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.EventInfo2 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.EventInfo3 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.EventInfo4 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.EventInfo5 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.EventInfo6 = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.EventInfo7 = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.CopperOreTier = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.IronOreTier = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.SilverOreTier = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.GoldOreTier = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.CobaltOreTier = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.MythrilOreTier = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.AdamantiteOreTier = tmp
	}
	{
		tmp, err := cur.ReadInt8()
		if err != nil {
			return err
		}
		m.InvasionType = tmp
	}
	{
		tmp, err := cur.ReadUint64()
		if err != nil {
			return err
		}
		m.LobbyID = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.SandstormSeverity = tmp
	}
	return nil
}