package api

import "github.com/patryklyczko/transport_app/pkg/db"

func (a *InstanceAPI) Orders() ([]db.Order, error) {
	return a.dbController.Orders()
}

func (a *InstanceAPI) Order(ID string) (*db.Order, error) {
	return a.dbController.Order(ID)
}

func (a *InstanceAPI) AddOrder(order *db.OrderRequest) (string, error) {
	return a.dbController.AddOrder(order)
}

func (a *InstanceAPI) UpdateOrder(order *db.Order) error {
	return a.dbController.UpdateOrder(order)
}

func (a *InstanceAPI) DeleteOrder(ID string) error {
	return a.dbController.DeleteOrder(ID)
}
