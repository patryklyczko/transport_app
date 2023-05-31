package algorithms

import "github.com/patryklyczko/transport_app/pkg/structures"

func Gain(solutions []structures.Solution) float32 {
	gain := float32(0)
	for _, solution := range solutions {
		orders := solution.Orders
		for _, order := range orders {
			gain += float32(order.Gain)
			gain -= solution.Driver.Capacity * 2
		}
	}
	return gain
}
