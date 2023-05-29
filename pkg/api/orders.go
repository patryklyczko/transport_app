package api

import "github.com/patryklyczko/transport_app/pkg/structures"

func (a *InstanceAPI) Orders() ([]structures.Order, error) {
	return a.dbController.Orders()
}

func (a *InstanceAPI) Order(ID string) (*structures.Order, error) {
	return a.dbController.Order(ID)
}

func (a *InstanceAPI) AddOrder(order *structures.OrderRequest) (string, error) {
	return a.dbController.AddOrder(order)
}

func (a *InstanceAPI) UpdateOrder(order *structures.Order) error {
	return a.dbController.UpdateOrder(order)
}

func (a *InstanceAPI) DeleteOrder(ID string) error {
	return a.dbController.DeleteOrder(ID)
}
