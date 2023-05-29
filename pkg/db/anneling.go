package db

import (
	"math"
	"math/rand"
	"time"

	"github.com/patryklyczko/transport_app/pkg/algorithms"
)

type AnnelingParameters struct {
	T_init  float32 `json:"t_init" bson:"t_init"`
	Cooling float32 `json:"cooling" bson:"cooling"`
	T_end   float32 `json:"t_end" bson:"t_end"`
	N_max   int32   `json:"n_max" bson:"n_max"`
	K       float64 `json:"k" bson:"k"`
}

type Solution struct {
	Driver     Driver
	Orders     []OrderAlgorithm
	EndTime    time.Time
	WeightLeft float32
}

func (d *DBController) Anneling(parameters *AnnelingParameters) (map[*Driver][]Order, float32, error) {
	var emptyOrders []Order
	var drivers []Driver
	var err error
	var freeOrders []Order
	var ordersPriority []OrderAlgorithm

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

	iteration := 0
	d.log.Debugf("Iteration \t Best_value \t Current value")
	actualNeighbor, freeOrders = d.ChangeNeighborhood(initialSolution, emptyOrders)
	for (temperature > parameters.T_end) && parameters.N_max > 0 {
		solutionNeighbor, freeOrders = d.ChangeNeighborhood(actualNeighbor, freeOrders)
		gainNeighbor := d.Gain(solutionNeighbor)
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
