// Code generated by gomacro; DO NOT EDIT.
package messages

import "github.com/tdakkota/cursor"

func (m NPCUpdate) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteInt16(m.NPCID)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.PositionX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteFloat32(m.PositionY)
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
	{
		err := cur.WriteUint16(m.Target)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(uint8(m.NpcFlags1))
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(uint8(m.NpcFlags2))
		if err != nil {
			return err
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI0 != 0 {
		{
			err := cur.WriteFloat32(m.AI0)
			if err != nil {
				return err
			}
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI1 != 0 {
		{
			err := cur.WriteFloat32(m.AI1)
			if err != nil {
				return err
			}
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI2 != 0 {
		{
			err := cur.WriteFloat32(m.AI2)
			if err != nil {
				return err
			}
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI3 != 0 {
		{
			err := cur.WriteFloat32(m.AI3)
			if err != nil {
				return err
			}
		}
	}
	{
		err := cur.WriteInt16(m.NPCNetID)
		if err != nil {
			return err
		}
	}
	if m.NpcFlags2&

		NPCUpdateNpcFlags2StatsScaled !=
		0 {
		{
			err := cur.WriteUint8(m.PlayerCountForMultiplayerDifficultyOverride)
			if err != nil {
				return err
			}
		}
	}
	if m.NpcFlags2&

		NPCUpdateNpcFlags2StrengthMultiplier !=
		0 {
		{
			err := cur.WriteFloat32(m.StrengthMultiplier)
			if err != nil {
				return err
			}
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1LifeMax !=
		0 {
		{
			err := cur.WriteUint8(m.LifeBytes)
			if err != nil {
				return err
			}
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1LifeMax !=
		0 {
		{
			err := cur.WriteInt32(m.Life)
			if err != nil {
				return err
			}
		}
	}
	{
		err := cur.WriteUint8(m.ReleaseOwner)
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *NPCUpdate) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.NPCID = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.PositionX = tmp
	}
	{
		tmp, err := cur.ReadFloat32()
		if err != nil {
			return err
		}
		m.PositionY = tmp
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
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.Target = tmp
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.NpcFlags1 = NPCUpdateNpcFlags1(tmp)
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.NpcFlags2 = NPCUpdateNpcFlags2(tmp)
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI0 != 0 {
		{
			tmp, err := cur.ReadFloat32()
			if err != nil {
				return err
			}
			m.AI0 = tmp
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI1 != 0 {
		{
			tmp, err := cur.ReadFloat32()
			if err != nil {
				return err
			}
			m.AI1 = tmp
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI2 != 0 {
		{
			tmp, err := cur.ReadFloat32()
			if err != nil {
				return err
			}
			m.AI2 = tmp
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1AI3 != 0 {
		{
			tmp, err := cur.ReadFloat32()
			if err != nil {
				return err
			}
			m.AI3 = tmp
		}
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.NPCNetID = tmp
	}
	if m.NpcFlags2&

		NPCUpdateNpcFlags2StatsScaled !=
		0 {
		{
			tmp, err := cur.ReadUint8()
			if err != nil {
				return err
			}
			m.PlayerCountForMultiplayerDifficultyOverride = tmp
		}
	}
	if m.NpcFlags2&

		NPCUpdateNpcFlags2StrengthMultiplier !=
		0 {
		{
			tmp, err := cur.ReadFloat32()
			if err != nil {
				return err
			}
			m.StrengthMultiplier = tmp
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1LifeMax !=
		0 {
		{
			tmp, err := cur.ReadUint8()
			if err != nil {
				return err
			}
			m.LifeBytes = tmp
		}
	}
	if m.NpcFlags1&

		NPCUpdateNpcFlags1LifeMax !=
		0 {
		{
			tmp, err := cur.ReadInt32()
			if err != nil {
				return err
			}
			m.Life = tmp
		}
	}
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.ReleaseOwner = tmp
	}
	return nil
}
