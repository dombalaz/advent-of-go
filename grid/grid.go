package grid

type Grid[T any] struct {
	x, y int
	data []T
}

func New[T any](x, y int) Grid[T] {
	return Grid[T]{x, y, make([]T, x*y)}
}

func NewFromSlice[T any](x, y int, vals []T) Grid[T] {
	return Grid[T]{x: x, y: y, data: append([]T{}, vals...)}
}

func (g *Grid[T]) Fill(vs []T) {
	copy(g.data, vs)
}

func (g *Grid[T]) At(x, y int) T {
	return g.data[g.i(x, y)]
}

func (g *Grid[T]) Set(x, y int, v T) {
	g.data[g.i(x, y)] = v
}

func (g *Grid[T]) X() int {
	return g.x
}

func (g *Grid[T]) Y() int {
	return g.y
}

func (g *Grid[T]) i(x, y int) int {
	return y*g.x + x
}
