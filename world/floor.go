package world

import "sync"

type Floor struct {
	Width, Height int
	Tiles         [][]Tile
	mu            sync.Mutex
}

func (f *Floor) GetTile(x, y int) *Tile {
	return &f.Tiles[y][x]
}
