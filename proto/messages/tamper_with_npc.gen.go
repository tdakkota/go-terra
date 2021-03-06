// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m TamperWithNPC) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint16(m.NPCID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteBool(m.SetNPCImmunity)
		if err != nil {
			return err
		}
	}
	if m.SetNPCImmunity {
		{
			err := cur.WriteInt32(m.ImmunityTime)
			if err != nil {
				return err
			}
		}
	}
	{
		err := cur.WriteInt16(m.ImmunityPlayerID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *TamperWithNPC) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.NPCID = tmp
	}
	{
		tmp, err := cur.ReadBool()
		if err != nil {
			return err
		}
		m.SetNPCImmunity = tmp
	}
	if m.SetNPCImmunity {
		{
			tmp, err := cur.ReadInt32()
			if err != nil {
				return err
			}
			m.ImmunityTime = tmp
		}
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ImmunityPlayerID = tmp
	}
	return nil
}
