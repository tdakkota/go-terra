// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m UpdateNPCBuff) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.NPCID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.BuffID1)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Time1)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.BuffID2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Time2)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.BuffID3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Time3)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.BuffID4)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Time4)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint16(m.BuffID5)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Time5)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *UpdateNPCBuff) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.NPCID = tmp
	}
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.BuffID1 = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Time1 = tmp
	}
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.BuffID2 = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Time2 = tmp
	}
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.BuffID3 = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Time3 = tmp
	}
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.BuffID4 = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Time4 = tmp
	}
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.BuffID5 = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Time5 = tmp
	}
	return nil
}
