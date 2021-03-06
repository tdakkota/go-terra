// Code generated by gomacro; DO NOT EDIT.
package structs

import "github.com/tdakkota/cursor"

func (m LiquidType) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(uint8(m))
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *LiquidType) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		*m = LiquidType(tmp)
	}
	return nil
}
