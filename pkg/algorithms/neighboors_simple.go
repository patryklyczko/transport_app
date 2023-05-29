package algorithms

import "github.com/patryklyczko/transport_app/pkg/db"

func NeighboorsSimpl(orders []db.OrderAlgorithm, drivers []db.Driver) []db.Solution {
	solutions := []db.Solution{}

	for i, driver := range drivers {
		solution := db.Solution{
			Driver:     driver,
			Orders:     []db.OrderAlgorithm{orders[i]},
			EndTime:    orders[i].TimeFinish,
			WeightLeft: driver.Capacity - orders[i].Weight,
		}
		orders[i].Taken = true
		solutions = append(solutions, solution)
	}
	return solutions
}
