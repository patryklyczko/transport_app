package db

import (
	"context"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func (d *DBController) Drivers() ([]structures.Driver, error) {
	collection := d.db.Collection("Drivers")
	var drivers []structures.Driver
	filter := bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var driver structures.Driver
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

func (d *DBController) Driver(ID string) (*structures.Driver, error) {
	collection := d.db.Collection("Drivers")
	var driver *structures.Driver

	filter := bson.M{"id": ID}
	if err := collection.FindOne(context.Background(), filter).Decode(&driver); err != nil {
		return nil, err
	}

	return driver, nil
}

func (d *DBController) PageDriver(page, number int) (*structures.DriverPagination, error) {
	collection := d.db.Collection("Drivers")
	var drivers []structures.Driver

	skip := page * number
	findOptions := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(number))
	filter := bson.M{}

	countOrder, _ := collection.CountDocuments(context.Background(), filter)

	cur, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var driver structures.Driver
		if err := cur.Decode(&driver); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &structures.DriverPagination{DriversCount: int(countOrder), Drivers: drivers}, nil
}

func (d *DBController) AddDriver(driverRequest *structures.DriverRequest) (string, error) {
	collection := d.db.Collection("Drivers")
	ID := generateID("dr")

	driver := structures.Driver{
		ID:       ID,
		Position: driverRequest.Position,
		Name:     driverRequest.Name,
		Orders:   driverRequest.Orders,
		Capacity: driverRequest.Capacity,
	}

	_, err := collection.InsertOne(context.Background(), driver)
	if err != nil {
		d.log.Debugf("error while inserting order")
	}

	return ID, nil
}

func (d *DBController) UpdateDriver(driver *structures.Driver) error {
	collection := d.db.Collection("Drivers")

	filter := bson.M{"id": driver.ID}
	update := bson.M{"$set": driver}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := collection.FindOneAndUpdate(context.Background(), filter, update, options)
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
