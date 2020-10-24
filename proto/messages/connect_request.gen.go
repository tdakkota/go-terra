// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m ConnectRequest) Append(cur *cursor.Cursor) (err error) {
	{
		_, err := cur.WriteString(m.Version)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *ConnectRequest) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadString()
		if err != nil {
			return err
		}
		m.Version = tmp
	}
	return nil
}
