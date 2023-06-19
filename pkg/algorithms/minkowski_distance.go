package algorithms

import (
	"github.com/patryklyczko/transport_app/pkg/structures"
)

func MinkowskiDistance(posStart, posEnd structures.Position) float32 {
	return abs(structures.LatScaler*(posStart.Lat-posEnd.Lat)) + abs(structures.LatScaler*(posStart.Lon-posEnd.Lon))
}

func abs(value float32) float32 {
	if value < 0 {
		return -value
	}
	return value
}

func TimeToCompleteOrder(order *structures.OrderAlgorithm, driverPos structures.Position) float32 {
	distOrder := MinkowskiDistance(order.PositionSend, order.PositionTake)
	distDriverOrder := MinkowskiDistance(driverPos, order.PositionTake)
	return (distDriverOrder + distOrder) / 60
}
