package db

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func generateID(name string) string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	randomNum := rand.Intn(10)
	uid := fmt.Sprintf("%s_%d-%d", name, now, randomNum)
	return uid
}

func (d *DBController) Orders() ([]structures.Order, error) {
	collection := d.db.Collection("Orders")
	var orders []structures.Order
	filter := bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var order structures.Order
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

func (d *DBController) Order(ID string) (*structures.Order, error) {
	collection := d.db.Collection("Orders")
	var order *structures.Order

	filter := bson.M{"id": ID}
	if err := collection.FindOne(context.Background(), filter).Decode(&order); err != nil {
		return nil, err
	}

	return order, nil
}

func (d *DBController) AddOrder(orderRequest *structures.OrderRequest) (string, error) {
	collection := d.db.Collection("Orders")
	ID := generateID("od")

	order := structures.Order{
		ID:           ID,
		PositionTake: orderRequest.PositionTake,
		PositionSend: orderRequest.PositionSend,
		TimeAdd:      orderRequest.TimeAdd,
		TimeEnd:      orderRequest.TimeEnd,
		Gain:         orderRequest.Gain,
	}

	_, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		d.log.Debugf("error while inserting order")
	}

	return ID, nil
}

func (d *DBController) UpdateOrder(order *structures.Order) error {
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

func (d *DBController) EmptyOrders() ([]structures.Order, error) {
	var orders []structures.Order
	collection := d.db.Collection("Orders")
	filter := bson.M{"taken": false}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var order structures.Order
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
