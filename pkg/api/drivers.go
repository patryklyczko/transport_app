package api

import "github.com/patryklyczko/transport_app/pkg/structures"

func (a *InstanceAPI) Drivers() ([]structures.Driver, error) {
	return a.dbController.Drivers()
}

func (a *InstanceAPI) Driver(ID string) (*structures.Driver, error) {
	return a.dbController.Driver(ID)
}

func (a *InstanceAPI) PageDriver(page, number int) (*structures.DriverPagination, error) {
	return a.dbController.PageDriver(page, number)
}

func (a *InstanceAPI) AddDriver(driver *structures.DriverRequest) (string, error) {
	return a.dbController.AddDriver(driver)
}

func (a *InstanceAPI) UpdateDriver(driver *structures.Driver) error {
	return a.dbController.UpdateDriver(driver)
}

func (a *InstanceAPI) DeleteDriver(ID string) error {
	return a.dbController.DeleteDriver(ID)
}
