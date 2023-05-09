package db

import (
	"math"
	"math/rand"
)

type AnnelingParameters struct {
	T_init  float32 `json:"t_init" bson:"t_init"`
	Cooling float32 `json:"cooling" bson:"cooling"`
	T_end   float32 `json:"t_end" bson:"t_end"`
	N_max   int32   `json:"n_max" bson:"n_max"`
	K       float64 `json:"k" bson:"k"`
}

func (d *DBController) Anneling(parameters *AnnelingParameters) (map[*Driver][]Order, float32, error) {
	var emptyOrders []Order
	var drivers []Driver
	var err error

	if emptyOrders, err = d.EmptyOrders(); err != nil {
		return nil, 0, err
	}
	if drivers, err = d.Drivers(); err != nil {
		return nil, 0, err
	}

	initialSolution, emptyOrders, err := d.NeighboorsSimple(emptyOrders, drivers)
	if err != nil {
		return nil, 0, err
	}

	bestSolution := initialSolution
	bestGain := d.Gain(bestSolution)
	temperature := parameters.T_init

	for (temperature > parameters.T_end) || parameters.N_max > 0 {
		solutionNeighbor := d.ChangeNeighborhood(initialSolution, emptyOrders)
		gainNeighbor := d.Gain(solutionNeighbor)
		if gainNeighbor < bestGain {
			bestSolution = solutionNeighbor
			bestGain = gainNeighbor
		} else {
			delta := gainNeighbor - bestGain
			r := rand.Float64()
			if r < math.Exp(-float64(delta)/(parameters.K*float64(temperature))) {
				bestSolution = solutionNeighbor
				bestGain = gainNeighbor
			}
		}
		temperature *= parameters.Cooling
	}

	return bestSolution, bestGain, nil
}
