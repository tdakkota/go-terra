// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m RemoveRevengeMarker) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt32(m.UniqueID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *RemoveRevengeMarker) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.UniqueID = tmp
	}
	return nil
}
