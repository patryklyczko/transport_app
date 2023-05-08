package db

import (
	"container/heap"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DBController) FindNearestNode(position *Position) (*NodePositions, error) {
	collection := d.db.Collection("Nodes")
	var cur *mongo.Cursor
	var node NodePositions
	var err error

	filter := bson.M{}
	options := options.Find()
	options.SetSort(bson.M{
		"$sqrt": bson.M{
			"$add": []interface{}{
				bson.M{"$pow": bson.A{bson.M{"$subtract": []interface{}{"$lat", position.Lat}}, 2}},
				bson.M{"$pow": bson.A{bson.M{"$subtract": []interface{}{"$lon", position.Lon}}, 2}},
			},
		},
	})
	options.SetLimit(1)
	if cur, err = collection.Find(context.Background(), filter, options); err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		err := cur.Decode(&node)
		if err != nil {
			return nil, err
		}
	}
	return &node, nil
}

func (d *DBController) NodeToRelation(node *NodePositions) (*NodesRelations, error) {
	collection := d.db.Collection("Relations")
	var nodeRelation *NodesRelations

	filter := bson.M{"parent": node.Parent}
	if err := collection.FindOne(context.Background(), filter).Decode(&nodeRelation); err != nil {
		return nil, err
	}

	return nodeRelation, nil
}

func abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func heuristic(start Position, end Position) time.Duration {
	return time.Duration((abs(start.Lat-end.Lat) + abs(start.Lon-end.Lon)) * float32(time.Hour))
}

func (d *DBController) AStar(driver Driver, start Position, end Position) time.Duration {
	var nodes []NodesRelations
	var node *NodePositions
	var nodeStart *NodesRelations
	var nodeEnd *NodesRelations
	var err error

	if nodes, err = d.MapNodes(); err != nil {
		return 0
	}

	// Start Node
	if node, err = d.FindNearestNode(&start); err != nil {
		return 0
	}
	if nodeStart, err = d.NodeToRelation(node); err != nil {
		return 0
	}

	// End Node
	if node, err = d.FindNearestNode(&start); err != nil {
		return 0
	}
	if nodeEnd, err = d.NodeToRelation(node); err != nil {
		return 0
	}

	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)
	heap.Push(&openSet, &Item{Node: nodeStart})

}
