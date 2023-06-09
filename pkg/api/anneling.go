package api

import (
	"github.com/patryklyczko/transport_app/pkg/db"
	"github.com/patryklyczko/transport_app/pkg/structures"
)

func (a *InstanceAPI) Anneling(parameters *db.AnnelingParameters) (*structures.SolutionValues, error) {
	return a.dbController.Anneling(parameters)
}
