package algorithms

import "github.com/patryklyczko/transport_app/pkg/structures"

func PriorityOrders(orders []structures.Order) *Stack {
	var orderAlgorithms Stack
	var priority float32
	KScalar := float32(1000)
	for _, order := range orders {
		priority = KScalar * float32(order.Gain) / (MinkowskiDistance(order.PositionTake, order.PositionSend))
		orderAlg := structures.OrderAlgorithm{
			ID:         order.ID,
			Gain:       order.Gain,
			TimeFinish: order.TimeFinish,
			Weight:     order.Weight,
			Priority:   priority,
		}
		orderAlgorithms.Push(orderAlg)
	}
	orderAlgorithms.SortByPriority()
	return &orderAlgorithms
}
