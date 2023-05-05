package api

import "github.com/patryklyczko/transport_app/pkg/db"

func (a *InstanceAPI) AddOrder(order *db.OrderRequest) (string, error) {
	return a.dbController.AddOrder(order)
}

func (a *InstanceAPI) UpdateOrder(order *db.Order) error {
	return a.dbController.UpdateOrder(order)
}

func (a *InstanceAPI) DeleteOrder(ID string) error {
	return a.dbController.DeleteOrder(ID)
}
