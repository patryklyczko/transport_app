package db

type AnnelingParameters struct {
	T_init  float32 `json:"t_init" bson:"t_init"`
	Cooling float32 `json:"cooling" bson:"cooling"`
	T_end   float32 `json:"t_end" bson:"t_end"`
	N_max   int32   `json:"n_max" bson:"n_max"`
}

// func CheckTime()

// func Neighboors()

func (d *DBController) Anneling(parameters *AnnelingParameters) error {
	return nil
}
