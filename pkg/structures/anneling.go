package structures

import "time"

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
	FreeTime   []time.Time
	EndTime    time.Time
	WeightLeft float32
}

type SolutionValues struct {
	Profit float32 `json:"profit"`
}
