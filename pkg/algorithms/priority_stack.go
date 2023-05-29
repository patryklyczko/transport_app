package algorithms

import (
	"sort"

	"github.com/patryklyczko/transport_app/pkg/db"
)

func (s *Stack) Push(order db.OrderAlgorithm) {
	s.items = append(s.items, order)
}

func (s *Stack) SortByPriority() {
	sort.Slice(s.items, func(i, j int) bool {
		return s.items[i].Priority < s.items[j].Priority
	})
}

type Stack struct {
	items []db.OrderAlgorithm
}
