// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m SendPassword) Append(cur *cursor.Cursor) (err error) {
	{
		_, err := cur.WriteString(m.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *SendPassword) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadString()
		if err != nil {
			return err
		}
		m.Password = tmp
	}
	return nil
}
