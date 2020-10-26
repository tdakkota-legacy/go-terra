// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m UpdateItemOwner) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.ItemID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *UpdateItemOwner) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ItemID = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerID = tmp
	}
	return nil
}