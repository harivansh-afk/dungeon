package world

import (
	"main/entity"
	"sync"
)

type Floor struct {
	Width, Height int
	Tiles         [][]Tile
	mu            sync.Mutex
}

func NewFloor(width, height int) *Floor {
	tiles := make([][]Tile, height)
	for y := range tiles {
		tiles[y] = make([]Tile, width)
		for x := range tiles[y] {
			tiles[y][x] = Tile{Type: Wall}
		}
	}
	return &Floor{
		Width:  width,
		Height: height,
		Tiles:  tiles,
	}
}

func (f *Floor) GetTile(x, y int) *Tile {
	if x < 0 || x >= f.Width || y < 0 || y >= f.Height {
		return nil
	}
	return &f.Tiles[y][x]
}

func (f *Floor) MoveEntity(e *entity.Entity, newX, newY int) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	target := f.GetTile(newX, newY)

	if target == nil || target.Type == Wall || target.Occupant != nil {
		return false
	}

	f.GetTile(e.X, e.Y).Occupant = nil

	target.Occupant = e

	e.X, e.Y = newX, newY

	return true
}
