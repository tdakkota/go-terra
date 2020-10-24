// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m SpawnPlayer) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(m.PlayerID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.SpawnX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.SpawnY)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.RespawnTimeRemaining)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(m.PlayerSpawnContext)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *SpawnPlayer) Scan(cur *cursor.Cursor) (err error) {
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
		m.SpawnX = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.SpawnY = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.RespawnTimeRemaining = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.PlayerSpawnContext = tmp
	}
	return nil
}
