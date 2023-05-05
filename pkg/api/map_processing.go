package api

import "github.com/patryklyczko/transport_app/pkg/db"

func (a *InstanceAPI) ProcessMap(path *db.MapRequest) error {
	return a.dbController.ProcessMap(path)
}
