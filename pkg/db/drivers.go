package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Driver struct {
	ID       string   `json:"id" bson:"id"`
	Name     string   `json:"name" bson:"name"`
	Position Position `json:"position" bson:"position"`
	Orders   []Order  `json:"orders" bson:"orders"`
}

type DriverRequest struct {
	Name     string   `json:"name" bson:"name"`
	Position Position `json:"position" bson:"position"`
	Orders   []Order  `json:"orders" bson:"orders"`
}

func (d *DBController) Drivers() ([]Driver, error) {
	collection := d.db.Collection("Drivers")
	var drivers []Driver
	filter := bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var driver Driver
		if err := cur.Decode(&driver); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return drivers, nil
}

func (d *DBController) Driver(ID string) (*Driver, error) {
	collection := d.db.Collection("Drivers")
	var driver *Driver

	filter := bson.M{"id": ID}
	if err := collection.FindOne(context.Background(), filter).Decode(&driver); err != nil {
		return nil, err
	}

	return driver, nil
}

func (d *DBController) AddDriver(driverRequest *DriverRequest) (string, error) {
	collection := d.db.Collection("Drivers")
	ID := generateID("dr")

	driver := Driver{
		ID:       ID,
		Position: driverRequest.Position,
		Name:     driverRequest.Name,
		Orders:   driverRequest.Orders,
	}

	_, err := collection.InsertOne(context.Background(), driver)
	if err != nil {
		d.log.Debugf("error while inserting order")
	}

	return ID, nil
}

func (d *DBController) UpdateDriver(driver *Driver) error {
	collection := d.db.Collection("Drivers")

	filter := bson.M{"id": driver.ID}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := collection.FindOneAndUpdate(context.Background(), filter, driver, options)
	if err := err.Err(); err != nil {
		return err
	}

	return nil
}

func (d *DBController) DeleteDriver(ID string) error {
	collection := d.db.Collection("Drivers")

	filter := bson.M{"id": ID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
