// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m SetNPCKillCount) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.NPCType)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.KillCount)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *SetNPCKillCount) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.NPCType = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.KillCount = tmp
	}
	return nil
}
