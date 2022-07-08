package state

import (
	"battleground/constants"
)

// Map contains ground, static object and some other infomation about ground
type Map struct {
	tiles [][]int
	// 0: no ground, user fall and die
	// 1: ground, user can walk, run ...
	// 2: wall, user cannot move to wall
	// 3: mud, user slow down
	// 4: boost, user move faster
	// 5: flow, user pushed to one direction

	// Access tiles[y][x]
}

func NewMap() *Map {
	return &Map{
		tiles: [][]int{
			{0, 0, 2, 2, 2, 2, 2, 0, 1},
			{0, 1, 1, 1, 1, 1, 2, 1, 1},
			{2, 1, 2, 1, 1, 1, 2, 0, 1},
			{2, 1, 1, 1, 1, 1, 2, 1, 1},
			{2, 1, 1, 1, 2, 1, 2, 2, 2},
			{2, 1, 1, 1, 1, 1, 2, 2, 2},
			{2, 2, 2, 2, 2, 0, 2, 0, 2},
		},
	}
}

func (m *Map) GetTilesAt(x, y int) int {
	if x < m.GetXn() && x >= 0 && y < m.GetYn() && y >= 0 {
		return m.tiles[y][x]
	}
	return -99
}

func (m *Map) GetYn() int {
	return len(m.tiles)
}

func (m *Map) GetXn() int {
	return len(m.tiles[0])
}

func (m *Map) TileContain(object Object) (int32, int32) {
	return object.Position().X / constants.MapTileSize, object.Position().Y / constants.MapTileSize
}
