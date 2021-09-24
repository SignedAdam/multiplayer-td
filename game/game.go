package game

type Game struct {
	World World
	Units []*Unit
}

type World struct {
	tiles [][]*Tile
}

type Unit struct {
	X int
	Y int
}

type Tile struct {
	Structure *Structure
}

type Structure struct {
}

func NewGame(size int) Game {
	return Game{
		World: World{
			tiles: make([][]*Tile, size, size),
		},
	}
}
