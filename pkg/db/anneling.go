package db

import (
	"math"
	"math/rand"

	"github.com/patryklyczko/transport_app/pkg/algorithms"
	"github.com/patryklyczko/transport_app/pkg/structures"
)

type AnnelingParameters struct {
	T_init  float32 `json:"t_init" bson:"t_init"`
	Cooling float32 `json:"cooling" bson:"cooling"`
	T_end   float32 `json:"t_end" bson:"t_end"`
	N_max   int32   `json:"n_max" bson:"n_max"`
	K       float64 `json:"k" bson:"k"`
}

func (d *DBController) Anneling(parameters *AnnelingParameters) ([]structures.Solution, float32, error) {
	var emptyOrders []structures.Order
	var drivers []structures.Driver
	var err error
	var ordersPriority *algorithms.Stack
	var solutionNeighbor []structures.Solution
	var actualNeighbor []structures.Solution

	if emptyOrders, err = d.EmptyOrders(); err != nil {
		return nil, 0, err
	}
	if drivers, err = d.Drivers(); err != nil {
		return nil, 0, err
	}
	ordersPriority = algorithms.PriorityOrders(emptyOrders)

	initialSolution := algorithms.NeighboorsSimple(ordersPriority, drivers)
	if err != nil {
		return nil, 0, err
	}

	bestSolution := initialSolution
	bestGain := algorithms.Gain(bestSolution)
	temperature := parameters.T_init

	if ordersPriority.IsEmpty() {
		return nil, 0, nil
	}

	iteration := 0
	d.log.Debugf("Iteration \t Best_value \t Current value")
	actualNeighbor = algorithms.ChangeNeighborhood(initialSolution, ordersPriority, 6)
	for (temperature > parameters.T_end) && parameters.N_max > 0 {
		solutionNeighbor = algorithms.ChangeNeighborhood(actualNeighbor, ordersPriority, 6)
		gainNeighbor := algorithms.Gain(solutionNeighbor)
		d.log.Debugf("%v \t\t %v \t\t %v", iteration, bestGain, gainNeighbor)

		if gainNeighbor > bestGain {
			bestSolution = solutionNeighbor
			bestGain = gainNeighbor
		}
		delta := gainNeighbor - bestGain
		r := rand.Float64()
		if r < math.Exp(-float64(delta)/(parameters.K*float64(temperature))) {
			actualNeighbor = solutionNeighbor
		} else {
			actualNeighbor = bestSolution
		}

		temperature *= parameters.Cooling
		parameters.N_max -= 1
		iteration += 1
	}

	return bestSolution, bestGain, nil
}
