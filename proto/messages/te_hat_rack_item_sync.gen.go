// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m TEHatRackItemSync) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.TileEntityID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.ItemIndex)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.ItemID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.Stack)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.Prefix)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *TEHatRackItemSync) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerID = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.TileEntityID = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.ItemIndex = tmp
	}
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.ItemID = tmp
	}
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.Stack = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Prefix = tmp
	}
	return nil
}
