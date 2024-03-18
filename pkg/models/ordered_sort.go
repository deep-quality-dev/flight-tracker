package models

// OrderedSet is a dynamic and insertion ordered set. When iterating over its elements,
// it will return them in order they were inserted originally.
type OrderedSet[T comparable] struct {
	indexes map[T]int
	items   []T
	length  int
}

// NewOrderedSet constructs a new instance of OrderedSet.
func NewOrderedSet[T comparable]() *OrderedSet[T] {
	return &OrderedSet[T]{
		indexes: make(map[T]int),
		length:  0,
	}
}

func (s *OrderedSet[T]) add(item T) bool {
	_, ok := s.indexes[item]
	if !ok {
		s.indexes[item] = s.length
		s.items = append(s.items, item)
		s.length++
	}

	return !ok
}

func (s *OrderedSet[T]) copy() *OrderedSet[T] {
	clone := NewOrderedSet[T]()
	for _, item := range s.items {
		clone.add(item)
	}

	return clone
}

func (s *OrderedSet[T]) indexOf(item T) int {
	idx, ok := s.indexes[item]
	if ok {
		return idx
	}

	return -1
}
