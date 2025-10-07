package grid

type Grid[T any] struct {
	x, y int
	data []T
}

func New[T any](x, y int) Grid[T] {
	return Grid[T]{x, y, make([]T, x*y)}
}

func (g *Grid[T]) Fill(vs []T) {
	copy(g.data, vs)
}

func (g *Grid[T]) At(x, y int) T {
	return g.data[y*g.x+x]
}

func (g *Grid[T]) X() int {
	return g.x
}

func (g *Grid[T]) Y() int {
	return g.y
}
