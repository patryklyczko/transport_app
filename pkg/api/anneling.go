package api

import "github.com/patryklyczko/transport_app/pkg/db"

func (a *InstanceAPI) Anneling(parameters *db.AnnelingParameters) error {
	return a.dbController.Anneling(parameters)
}
