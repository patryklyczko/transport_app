package api

import "github.com/patryklyczko/transport_app/pkg/db"

func (a *InstanceAPI) Anneling(parameters *db.AnnelingParameters) (map[*db.Driver][]db.Order, float32, error) {
	return a.dbController.Anneling(parameters)
}
