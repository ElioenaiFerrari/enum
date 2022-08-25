package enum

type Enum[T comparable] struct {
	Values []T
}

// New create new enum with values
// e := enum.New(0, 1, 2)
func New[T comparable](args ...T) *Enum[T] {
	return &Enum[T]{
		Values: args,
	}
}

// IsValid validate if arg is valid in enum
// e := enum.New(0, 1, 2)
// e.IsValid(0) => true
// e.IsValid(1) => true
// e.IsValid(2) => true
// e.IsValid(3) => false
func (e Enum[T]) IsValid(arg T) bool {
	for _, v := range e.Values {
		if v == arg {
			return true
		}
	}

	return false
}
