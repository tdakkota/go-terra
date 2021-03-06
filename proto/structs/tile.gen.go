// Code generated by gomacro; DO NOT EDIT.
package structs

import "github.com/tdakkota/cursor"

func (m TileFlags1) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(uint8(m))
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *TileFlags1) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		*m = TileFlags1(tmp)
	}
	return nil
}
func (m TileFlags2) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint8(uint8(m))
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *TileFlags2) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		*m = TileFlags2(tmp)
	}
	return nil
}
func (m Tile) Append(cur *cursor.Cursor) (err error) {
	{
		err := m.Flags1.Append(cur)
		if err != nil {
			return err
		}
	}
	{
		err := m.Flags2.Append(cur)
		if err != nil {
			return err
		}
	}
	if m.Flags2&HasColor !=
		0 {
		{
			err := cur.WriteUint8(m.Color)
			if err != nil {
				return err
			}
		}
	}
	if m.Flags2&HasWallColor !=
		0 {
		{
			err := cur.WriteUint8(m.WallColor)
			if err != nil {
				return err
			}
		}
	}
	if m.Flags1&Active !=
		0 {
		{
			err := cur.WriteUint16(m.Type)
			if err != nil {
				return err
			}
		}
	}
	if m.Flags1&Active !=
		0 {
		{
			err := cur.WriteInt16(m.FrameX)
			if err != nil {
				return err
			}
		}
	}
	if m.Flags1&Active !=
		0 {
		{
			err := cur.WriteInt16(m.FrameY)
			if err != nil {
				return err
			}
		}
	}
	if m.Flags1&HasWall !=
		0 {
		{
			err := cur.WriteUint16(m.Wall)
			if err != nil {
				return err
			}
		}
	}
	if m.Flags1&HasLiquid !=
		0 {
		{
			err := cur.WriteUint8(m.Liquid)
			if err != nil {
				return err
			}
		}
	}
	if m.Flags1&HasLiquid !=
		0 {
		{
			err := m.LiquidType.Append(cur)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (m *Tile) Scan(cur *cursor.Cursor) (err error) {
	{
		err := m.Flags1.Scan(cur)
		if err != nil {
			return err
		}
	}
	{
		err := m.Flags2.Scan(cur)
		if err != nil {
			return err
		}
	}
	if m.Flags2&HasColor !=
		0 {
		{
			tmp, err := cur.ReadUint8()
			if err != nil {
				return err
			}
			m.Color = tmp
		}
	}
	if m.Flags2&HasWallColor !=
		0 {
		{
			tmp, err := cur.ReadUint8()
			if err != nil {
				return err
			}
			m.WallColor = tmp
		}
	}
	if m.Flags1&Active !=
		0 {
		{
			tmp, err := cur.ReadUint16()
			if err != nil {
				return err
			}
			m.Type = tmp
		}
	}
	if m.Flags1&Active !=
		0 {
		{
			tmp, err := cur.ReadInt16()
			if err != nil {
				return err
			}
			m.FrameX = tmp
		}
	}
	if m.Flags1&Active !=
		0 {
		{
			tmp, err := cur.ReadInt16()
			if err != nil {
				return err
			}
			m.FrameY = tmp
		}
	}
	if m.Flags1&HasWall !=
		0 {
		{
			tmp, err := cur.ReadUint16()
			if err != nil {
				return err
			}
			m.Wall = tmp
		}
	}
	if m.Flags1&HasLiquid !=
		0 {
		{
			tmp, err := cur.ReadUint8()
			if err != nil {
				return err
			}
			m.Liquid = tmp
		}
	}
	if m.Flags1&HasLiquid !=
		0 {
		{
			err := m.LiquidType.Scan(cur)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
