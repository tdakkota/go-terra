// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m MassWireOperationConsume) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.ItemType)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Quantity)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *MassWireOperationConsume) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ItemType = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Quantity = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerID = tmp
	}
	return nil
}
