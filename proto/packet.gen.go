// Code generated by gomacro; DO NOT EDIT.
package proto

import "github.com/tdakkota/cursor"

func (m Packet) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint16(m.Length)
		if err != nil {
			return err
		}
	}
	{
		err := m.Tag.Append(cur)
		if err != nil {
			return err
		}
	}
	{
		for _, v := range m.Message {
			{
				err := cur.WriteUint8(v)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (m *Packet) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.Length = tmp
	}
	{
		err := m.Tag.Scan(cur)
		if err != nil {
			return err
		}
	}
	{
		n := m.Length - 2
		m.Message = make([]byte, n)
		for i := 0; i < int(n); i++ {
			{
				tmp, err := cur.ReadUint8()
				if err != nil {
					return err
				}
				m.Message[i] = tmp
			}
		}
	}
	return nil
}
