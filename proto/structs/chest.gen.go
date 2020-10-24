// Code generated by gomacro; DO NOT EDIT.
package structs

import "github.com/tdakkota/cursor"

func (m Chest) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.Index)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.X)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Y)
		if err != nil {
			return err
		}
	}
	{
		_, err := cur.WriteString(m.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *Chest) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Index = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.X = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Y = tmp
	}
	{
		tmp, err := cur.ReadString()
		if err != nil {
			return err
		}
		m.Name = tmp
	}
	return nil
}
