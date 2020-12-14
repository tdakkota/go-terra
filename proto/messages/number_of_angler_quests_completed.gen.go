// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m NumberOfAnglerQuestsCompleted) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.AnglerQuestsCompleted)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.GolferScore)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *NumberOfAnglerQuestsCompleted) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerID = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.AnglerQuestsCompleted = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.GolferScore = tmp
	}
	return nil
}