// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m KillPortal) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint16(m.ProjectileOwner)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.ProjectileAI)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *KillPortal) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.ProjectileOwner = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.ProjectileAI = tmp
	}
	return nil
}
