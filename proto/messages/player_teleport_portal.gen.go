// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m PlayerTeleportPortal) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.PortalColorIndex)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.NewPositionX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.NewPositionY)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.VelocityX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.VelocityY)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *PlayerTeleportPortal) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerID = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.PortalColorIndex = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.NewPositionX = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.NewPositionY = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.VelocityX = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.VelocityY = tmp
	}
	return nil
}
