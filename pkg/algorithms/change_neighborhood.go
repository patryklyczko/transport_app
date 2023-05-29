package algorithms

import (
	"math/rand"

	"github.com/patryklyczko/transport_app/pkg/structures"
)

func ChangeNeighborhood(solutions []structures.Solution, ordersPriority *Stack, removeProbability float64) []structures.Solution {
	order := ordersPriority.Pop()
	for _, solution := range solutions {

		r := rand.Float64()
		if r > removeProbability {
			orderRemoved := solution.Orders[0]
			solution.Orders = solution.Orders[1:]
			ordersPriority.Push(orderRemoved)
		}

		if solution.EndTime.Minute() < order.TimeFinish.Minute() && solution.WeightLeft < order.Weight {
			continue
		}
		solution.Orders = append(solution.Orders, order)
		order = ordersPriority.Pop()
	}
	ordersPriority.Push(order)
	return solutions
}
