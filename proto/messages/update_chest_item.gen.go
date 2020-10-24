// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m UpdateChestItem) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.ChestID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.ItemSlot)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Stack)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Prefix)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.ItemNetID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *UpdateChestItem) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ChestID = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.ItemSlot = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Stack = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Prefix = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ItemNetID = tmp
	}
	return nil
}
