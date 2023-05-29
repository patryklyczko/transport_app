package db

import (
	"context"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *DBController) MapNodes() ([]structures.NodesRelations, error) {
	collection := d.db.Collection("Relations")
	var nodes []structures.NodesRelations
	filter := bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var node structures.NodesRelations
		if err := cur.Decode(&node); err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return nodes, nil
}
