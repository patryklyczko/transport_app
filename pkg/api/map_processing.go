package api

import "github.com/patryklyczko/transport_app/pkg/structures"

func (a *InstanceAPI) ProcessMap(path *structures.MapRequest) error {
	return a.dbController.ProcessMap(path)
}
