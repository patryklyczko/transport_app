package algorithms

import "github.com/patryklyczko/transport_app/pkg/db"

func MinkowskiDistance(posStart, posEnd db.Position) float32 {
	return abs(posStart.Lat-posEnd.Lat) + abs(posStart.Lon-posEnd.Lon)
}

func abs(value float32) float32 {
	if value < 0 {
		return -value
	}
	return value
}
