package models

// Node represents a graph node.
type Node[T comparable] map[T]bool

func (n Node[T]) connectsTo(key T) {
	n[key] = true
}

func (n Node[T]) edges() []T {
	var keys []T
	for k := range n {
		keys = append(keys, k)
	}

	return keys
}
