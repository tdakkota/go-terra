// Code generated by gomacro; DO NOT EDIT.
package messages

import (
	"github.com/tdakkota/cursor"
	"github.com/tdakkota/go-terra/proto/structs"
)

func (m SendSection) Append(cur *cursor.Cursor) (err error) {
	{
		err := cur.WriteBool(m.Compressed)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.XStart)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt32(m.YStart)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Width)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteInt16(m.Height)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(uint8(len(m.Tiles)))
		if err != nil {
			return err
		}
		for _, v := range m.Tiles {
			{
				err := v.Append(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	{
		err := cur.WriteInt16(m.ChestCount)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(uint8(len(m.Chests)))
		if err != nil {
			return err
		}
		for _, v := range m.Chests {
			{
				err := v.Append(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	{
		err := cur.WriteInt16(m.SignCount)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(uint8(len(m.Signs)))
		if err != nil {
			return err
		}
		for _, v := range m.Signs {
			{
				err := v.Append(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	{
		err := cur.WriteInt16(m.TileEntityCount)
		if err != nil {
			return err
		}
	}
	{
		err := cur.WriteUint8(uint8(len(m.TileEntities)))
		if err != nil {
			return err
		}
		for _, v := range m.TileEntities {
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
func (m *SendSection) Scan(cur *cursor.Cursor) (err error) {
	{
		tmp, err := cur.ReadBool()
		if err != nil {
			return err
		}
		m.Compressed = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.XStart = tmp
	}
	{
		tmp, err := cur.ReadInt32()
		if err != nil {
			return err
		}
		m.YStart = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Width = tmp
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.Height = tmp
	}
	{
		n, err := cur.ReadUint8()
		if err != nil {
			return err
		}
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
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.ChestCount = tmp
	}
	{
		n, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Chests = make([]structs.Chest, n)
		for i := 0; i < int(n); i++ {
			{
				err := m.Chests[i].Scan(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.SignCount = tmp
	}
	{
		n, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.Signs = make([]structs.Sign, n)
		for i := 0; i < int(n); i++ {
			{
				err := m.Signs[i].Scan(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	{
		tmp, err := cur.ReadInt16()
		if err != nil {
			return err
		}
		m.TileEntityCount = tmp
	}
	{
		n, err := cur.ReadUint8()
		if err != nil {
			return err
		}
		m.TileEntities = make([]structs.TileEntity, n)
		for i := 0; i < int(n); i++ {
			{
				err := m.TileEntities[i].Scan(cur)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}