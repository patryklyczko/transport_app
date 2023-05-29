package algorithms

import "github.com/patryklyczko/transport_app/pkg/db"

func PriorityOrders(orders []db.Order) []db.OrderAlgorithm {
	var orderAlgorithms Stack
	var priority float32
	KScalar := float32(1000)
	for _, order := range orders {
		priority = KScalar * float32(order.Gain) / (MinkowskiDistance(order.PositionTake, order.PositionSend))
		orderAlg := db.OrderAlgorithm{
			ID:         order.ID,
			Gain:       order.Gain,
			TimeFinish: order.TimeFinish,
			Weight:     order.Weight,
			Taken:      order.Taken,
			Priority:   priority,
		}
		orderAlgorithms.Push(orderAlg)
	}
	orderAlgorithms.SortByPriority()
	return orderAlgorithms.items
}
