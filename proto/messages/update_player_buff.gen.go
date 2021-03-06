// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m UpdatePlayerBuff) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	{
		for _, v := range m.BuffType {
			{
				err := cur.WriteUint16(v)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (m *UpdatePlayerBuff) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerID = tmp
	}
	{
		n := 22
		for i := 0; i < int(n); i++ {
			{
				tmp, err := cur.ReadUint16()
				if err != nil {
					return err
				}
				m.BuffType[i] = tmp
			}
		}
	}
	return nil
}
