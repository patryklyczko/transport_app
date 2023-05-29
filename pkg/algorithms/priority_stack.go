package algorithms

import (
	"sort"

	"github.com/patryklyczko/transport_app/pkg/structures"
)

func (s *Stack) Push(order structures.OrderAlgorithm) {
	s.items = append(s.items, order)
}

func (s *Stack) SortByPriority() {
	sort.Slice(s.items, func(i, j int) bool {
		return s.items[i].Priority < s.items[j].Priority
	})
}

func (s *Stack) Pop() structures.OrderAlgorithm {
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

type Stack struct {
	items []structures.OrderAlgorithm
}
