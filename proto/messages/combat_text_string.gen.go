// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m CombatTextString) Append(cur *cursor.Cursor) (err error) {
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
		err := m.CombatText.Append(cur)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *CombatTextString) Scan(cur *cursor.Cursor) (err error) {
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
		err := m.CombatText.Scan(cur)
		if err != nil {
			return err
		}
	}
	return nil
}