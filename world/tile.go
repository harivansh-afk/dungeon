package world

import "main/entity"

type TileType int

const (
	Wall TileType = iota
	Floor
	Door
	StairsDown
)

type Tile struct {
	Type     TileType
	Occupant *entity.Entity
}

func (t *Tile) Symbol() rune {
	if t.Occupant != nil {
		return t.Occupant.Symbol
	}
	switch t.Type {
	case Wall:
		return '#'
	case Floor:
		return '.'
	case Door:
		return '+'
	case StairsDown:
		return '>'
	default:
		return '?'
	}
}
