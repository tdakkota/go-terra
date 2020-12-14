// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m UpdateGoodEvil) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.Good)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Evil)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Crimson)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *UpdateGoodEvil) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Good = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Evil = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Crimson = tmp
	}
	return nil
}