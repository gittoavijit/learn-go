//  Do The sum of two variables taken as input

package twosums

type OpInput interface {
	int | float64
}

type Operator[T OpInput] struct {
	A, B T
}

func Sum[T OpInput](i, j T) T {
	return (i + j)
}

func NewOperator[T OpInput](i, j T) *Operator[T] {
	return &Operator[T]{
		A: i, B: j,
	}
}

func (o *Operator[T]) Sum() T {
	return (o.A + o.B)
}

func (o *Operator[T]) Difference() T {
	return (o.B - o.A)
}

func (o *Operator[T]) SumWith(i T) T {
	return o.Sum() + i
}
