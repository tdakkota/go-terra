// Code generated by gomacro; DO NOT EDIT.
package messages

import (
	"github.com/tdakkota/cursor"
	"github.com/tdakkota/go-terra/proto/structs"
)

func (m SendTileSquare) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteUint16(m.Size)
		if err != nil {
			return err
		}
	}
	if m.Size&0x8000 !=
		0 {
		{
			err := cur.WriteUint8(m.TileChangeType)
			if err != nil {
				return err
			}
		}
	}
	{
		err := cur.WriteInt16(m.TileX)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.TileY)
		if err != nil {
			return err
		}
	}
	{
		for _, v := range m.Tiles {
			{
				err := v.Append(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (m *SendTileSquare) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadUint16()
		if err != nil {
			return err
		}
		m.Size = tmp
	}
	if m.Size&0x8000 !=
		0 {
		{
			tmp, err := cur.ReadUint8()
			if err != nil {
				return err
			}
			m.TileChangeType = tmp
		}
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.TileX = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.TileY = tmp
	}
	{
		n := m.Size
		m.Tiles = make([]structs.Tile, n)
		for i := 0; i < int(n); i++ {
			{
				err := m.Tiles[i].Scan(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}