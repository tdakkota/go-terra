// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m SyncActiveChest) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.ChestID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.ChestX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.ChestY)
		if err != nil {
			return err
		}
	}
	{
		_, err := cur.WriteString(m.ChestName)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *SyncActiveChest) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ChestID = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ChestX = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ChestY = tmp
	}
	{
		tmp, err := cur.ReadString()
		if err != nil {
			return err
		}
		m.ChestName = tmp
	}
	return nil
}
