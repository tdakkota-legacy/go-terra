// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m PlaceObject) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.X)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Y)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Type)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Style)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Alternate)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt8(m.Random)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteBool(m.Direction)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *PlaceObject) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.X = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Y = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Type = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Style = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Alternate = tmp
	}
	{
		tmp, err := cur.ReadInt8()
		if err != nil {
			return err
		}
		m.Random = tmp
	}
	{
		tmp, err := cur.ReadBool()
		if err != nil {
			return err
		}
		m.Direction = tmp
	}
	return nil
}