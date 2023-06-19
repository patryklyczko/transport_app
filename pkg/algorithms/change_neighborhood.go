package algorithms

import (
	"math/rand"

	"github.com/patryklyczko/transport_app/pkg/structures"
)

func ChangeNeighborhood(solutions []structures.Solution, ordersPriority *Stack, removeProbability float64) []structures.Solution {

	var order *structures.OrderAlgorithm
	for i, solution := range solutions {
		if order = ordersPriority.TakeFree(); order != nil {
			if solution.EndTime.Minute() < order.TimeFinish.Minute() && solution.WeightLeft < order.Weight {
				continue
			}
			solutions[i].Orders = append(solution.Orders, *order)
			solutions[i].EndTime = order.TimeFinish
		}

		r := rand.Float64() * 10
		if r > removeProbability && len(solution.Orders) > 0 {
			orderRemoved := solution.Orders[0]
			solutions[i].Orders = solution.Orders[1:]
			ordersPriority.Freed(&orderRemoved)
		}
	}
	return solutions
}

// var order *structures.OrderAlgorithm
// for i, solution := range solutions {
// 	if len(solution.Orders)-1 < 0 {
// 		continue
// 	}
// 	lastOrder := solution.Orders[len(solution.Orders)-1]
// 	if order = ordersPriority.TakeFree(); order != nil {
// 		timeToCompleteNewRoute := TimeToCompleteOrder(order, lastOrder.PositionSend)
// 		if float32(solution.EndTime.Hour())+timeToCompleteNewRoute > float32(order.TimeFinish.Hour()) {
// 			continue
// 		}
// 		solutions[i].Orders = append(solutions[i].Orders, *order)
// 		// fmt.Printf("%v", solutions[i].Orders)
// 		solutions[i].EndTime = solutions[i].EndTime.Add(time.Duration(timeToCompleteNewRoute))
// 		solutions[i].FreeTime = append(solutions[i].FreeTime, solutions[i].EndTime)
// 	}

// 	r := rand.Float32()
// 	timesComplete := solution.FreeTime
// 	if len(timesComplete) < 1 {
// 		continue
// 	}
// 	for j := 1; j < len(timesComplete); j++ {
// 		fmt.Printf("values: %v", float32(timesComplete[j].Sub(timesComplete[j-1]).Minutes()))
// 		if float32(timesComplete[j].Sub(timesComplete[j-1]).Minutes()) > r {
// 			solutions[i].Orders = solutions[i].Orders[:j-1]
// 			solutions[i].EndTime = solutions[i].FreeTime[j-1]
// 			solutions[i].FreeTime = solutions[i].FreeTime[:j-1]
// 			for k := j - 1; k < len(timesComplete); k++ {
// 				ordersPriority.Push(solution.Orders[k])
// 			}

// 			continue
// 		}
// 	}

// }
// return solutions
