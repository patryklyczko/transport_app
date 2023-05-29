package algorithms

import "github.com/patryklyczko/transport_app/pkg/structures"

func MinkowskiDistance(posStart, posEnd structures.Position) float32 {
	return abs(posStart.Lat-posEnd.Lat) + abs(posStart.Lon-posEnd.Lon)
}

func abs(value float32) float32 {
	if value < 0 {
		return -value
	}
	return value
}
