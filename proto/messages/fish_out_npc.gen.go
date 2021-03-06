// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m FishOutNPC) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint16(m.X)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.Y)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.NPCID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *FishOutNPC) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.X = tmp
	}
	{
		tmp, err := cur.ReadUint16()
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
		m.NPCID = tmp
	}
	return nil
}
