package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Position struct {
	Lat float32 `json:"lat" bson:"lat"`
	Lon float32 `json:"lon" bson:"lon"`
}

type Order struct {
	ID       string    `json:"id" bson:"id"`
	Position Position  `json:"position" bson:"position"`
	TimeAdd  time.Time `json:"time_add" bson:"time_add"`
	TimeEnd  time.Time `json:"time_end" bson:"time_end"`
	Gain     int64     `json:"gain" bson:"gain"`
	// DriversAssign []Drivers
}

type OrderRequest struct {
	Position Position  `json:"position" bson:"position"`
	TimeAdd  time.Time `json:"time_add" bson:"time_add"`
	TimeEnd  time.Time `json:"time_end" bson:"time_end"`
	Gain     int64     `json:"gain" bson:"gain"`
}

func generateID() string {
	return ""
}

func (d *DBController) AddOrder(orderRequest *OrderRequest) (string, error) {
	collection := d.db.Collection("Orders")
	ID := generateID()

	order := Order{
		ID:       ID,
		Position: orderRequest.Position,
		TimeAdd:  orderRequest.TimeAdd,
		TimeEnd:  orderRequest.TimeEnd,
		Gain:     orderRequest.Gain,
	}

	_, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		d.log.Debugf("error while inserting order")
	}

	return ID, nil
}

func (d *DBController) UpdateOrder(order *Order) error {
	collection := d.db.Collection("Orders")

	filter := bson.M{"id": order.ID}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := collection.FindOneAndUpdate(context.Background(), filter, order, options)
	if err := err.Err(); err != nil {
		return err
	}

	return nil
}

func (d *DBController) DeleteOrder(ID string) error {
	collection := d.db.Collection("Orders")

	filter := bson.M{"_id": ID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
