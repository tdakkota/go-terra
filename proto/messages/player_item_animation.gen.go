// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m PlayerItemAnimation) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.ItemRotation)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.ItemAnimation)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *PlayerItemAnimation) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerID = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.ItemRotation = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ItemAnimation = tmp
	}
	return nil
}