// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m Time) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteBool(m.DayTime)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.TimeValue)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.SunModY)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.MoonModY)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *Time) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadBool()
		if err != nil {
			return err
		}
		m.DayTime = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.TimeValue = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.SunModY = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.MoonModY = tmp
	}
	return nil
}
