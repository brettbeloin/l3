package avl

type Gen interface {
	~int | ~string | ~float64
}

type node[T Gen] struct {
}

type Avl[T Gen] struct {
}

// Create a new Avl Tree
func NewTree[T Gen]() Avl[T] {
	return Avl[T]{}
}
