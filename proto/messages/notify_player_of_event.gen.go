// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m NotifyPlayerOfEvent) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.EventID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *NotifyPlayerOfEvent) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.EventID = tmp
	}
	return nil
}
