package domain

type Nullable[T any] struct {
	Set   bool
	Value *T
}
