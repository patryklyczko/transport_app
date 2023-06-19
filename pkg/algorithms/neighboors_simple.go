package algorithms

import (
	"time"

	"github.com/patryklyczko/transport_app/pkg/structures"
)

func NeighboorsSimple(ordersPriority *Stack, drivers []structures.Driver) []structures.Solution {
	solutions := []structures.Solution{}

	for _, driver := range drivers {
		order := ordersPriority.Pop()
		solution := structures.Solution{
			Driver:     driver,
			Orders:     []structures.OrderAlgorithm{order},
			FreeTime:   []time.Time{order.TimeFinish},
			EndTime:    order.TimeFinish, //Change
			WeightLeft: driver.Capacity - order.Weight,
		}
		solutions = append(solutions, solution)
	}
	return solutions
}
