package frozenlist

import (
	"github.com/craiggwilson/teakwood/util/collections/list"
	"github.com/craiggwilson/teakwood/util/iter"
)

var _ list.ReadOnly[int] = (*Frozen[int])(nil)

func NewFrozen[T any](l list.ReadOnly[T]) *Frozen[T] {
	return &Frozen[T]{l}
}

type Frozen[T any] struct {
	l list.ReadOnly[T]
}

func (l *Frozen[T]) Iter() iter.Iter[T] {
	return l.l.Iter()
}

func (l *Frozen[T]) Len() int {
	return l.l.Len()
}

func (l *Frozen[T]) Value(idx int) T {
	return l.l.Value(idx)
}