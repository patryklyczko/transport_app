package algorithms

import (
	"math/rand"

	"github.com/patryklyczko/transport_app/pkg/structures"
)

func ChangeNeighborhood(solutions []structures.Solution, ordersPriority *Stack, removeProbability float64) []structures.Solution {
	var order *structures.OrderAlgorithm
	for i, solution := range solutions {
		if order = ordersPriority.TakeFree(); order != nil {
			if solution.EndTime.Minute() < order.TimeFinish.Minute() && solution.WeightLeft < order.Weight {
				continue
			}
			solutions[i].Orders = append(solution.Orders, *order)
		}

		r := rand.Float64() * 10
		if r > removeProbability && len(solution.Orders) > 0 {
			orderRemoved := solution.Orders[0]
			solutions[i].Orders = solution.Orders[1:]
			ordersPriority.Freed(&orderRemoved)
		}
	}
	return solutions
}
