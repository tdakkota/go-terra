// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m TeleportationPotion) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.Type)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *TeleportationPotion) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Type = tmp
	}
	return nil
}
