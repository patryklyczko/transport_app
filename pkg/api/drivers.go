package api

import "github.com/patryklyczko/transport_app/pkg/db"

func (a *InstanceAPI) Drivers() ([]db.Driver, error) {
	return a.dbController.Drivers()
}

func (a *InstanceAPI) Driver(ID string) (*db.Driver, error) {
	return a.dbController.Driver(ID)
}

func (a *InstanceAPI) AddDriver(driver *db.DriverRequest) (string, error) {
	return a.dbController.AddDriver(driver)
}

func (a *InstanceAPI) UpdateDriver(driver *db.Driver) error {
	return a.dbController.UpdateDriver(driver)
}

func (a *InstanceAPI) DeleteDriver(ID string) error {
	return a.dbController.DeleteDriver(ID)
}
