package algorithms

import "github.com/patryklyczko/transport_app/pkg/structures"

func NeighboorsSimpl(ordersPriority *Stack, drivers []structures.Driver) []structures.Solution {
	solutions := []structures.Solution{}

	for _, driver := range drivers {
		order := ordersPriority.Pop()
		solution := structures.Solution{
			Driver:     driver,
			Orders:     []structures.OrderAlgorithm{order},
			EndTime:    order.TimeFinish,
			WeightLeft: driver.Capacity - order.Weight,
		}
		solutions = append(solutions, solution)
	}
	return solutions
}
