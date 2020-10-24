// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m CreateCombatText) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteFloat32(m.X)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.Y)
		if err != nil {
			return err
		}
	}
	{
		err := m.Color.Append(cur)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.HealAmount)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *CreateCombatText) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.X = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.Y = tmp
	}
	{
		err := m.Color.Scan(cur)
		if err != nil {
			return err
		}
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.HealAmount = tmp
	}
	return nil
}
