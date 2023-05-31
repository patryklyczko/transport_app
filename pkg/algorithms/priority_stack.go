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

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Get() []structures.OrderAlgorithm {
	return s.items
}

func (s *Stack) TakeFree() *structures.OrderAlgorithm {
	for i, order := range s.items {
		if !order.Taken {
			s.items[i].Taken = true
			return &s.items[i]
		}
	}
	return nil
}

func (s *Stack) Freed(orderFreed *structures.OrderAlgorithm) {
	for i, order := range s.items {
		if order.ID == orderFreed.ID {
			s.items[i].Taken = false
		}
	}
}

type Stack struct {
	items []structures.OrderAlgorithm
}
