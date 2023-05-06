package db

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type UID struct {
	ID string `json:"id" bson:"id"`
}

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

func generateID(name string) string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	randomNum := rand.Intn(10)
	uid := fmt.Sprintf("%s_%d-%d", name, now, randomNum)
	return uid
}

func (d *DBController) Orders() ([]Order, error) {
	collection := d.db.Collection("Orders")
	var orders []Order
	filter := bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var order Order
		if err := cur.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (d *DBController) Order(ID string) (*Order, error) {
	collection := d.db.Collection("Orders")
	var order *Order

	filter := bson.M{"id": ID}
	if err := collection.FindOne(context.Background(), filter).Decode(&order); err != nil {
		return nil, err
	}

	return order, nil
}

func (d *DBController) AddOrder(orderRequest *OrderRequest) (string, error) {
	collection := d.db.Collection("Orders")
	ID := generateID("od")

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

	filter := bson.M{"id": ID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
