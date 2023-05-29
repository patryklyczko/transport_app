package db

import (
	"container/heap"
	"context"
	"math"
	"time"

	"github.com/patryklyczko/transport_app/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DBController) FindNearestNode(position *structures.Position) (*structures.NodePositions, error) {
	collection := d.db.Collection("Nodes")
	var cur *mongo.Cursor
	var node structures.NodePositions
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

func (d *DBController) NodeToRelation(node *structures.NodePositions) (*structures.NodesRelations, error) {
	collection := d.db.Collection("Relations")
	var nodeRelation *structures.NodesRelations

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

func heuristic(start structures.Position, end structures.Position) time.Duration {
	return time.Duration((abs(start.Lat-end.Lat) + abs(start.Lon-end.Lon)) * float32(time.Hour))
}

func (d *DBController) AStar(driver structures.Driver, start structures.Position, end structures.Position) time.Duration {
	collection := d.db.Collection("Relations")
	collectionNodes := d.db.Collection("Nodes")
	var nodes []structures.NodesRelations
	var node *structures.NodePositions
	var nodeStart *structures.NodesRelations
	var nodeEnd *structures.NodesRelations
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
	heap.Push(&openSet, &Item{Node: nodeStart, Priority: float32(heuristic(start, end))})

	gScore := make(map[*structures.NodesRelations]time.Duration, len(nodes))
	fScore := make(map[*structures.NodesRelations]time.Duration, len(nodes))
	for _, node := range nodes {
		gScore[&node] = math.MaxInt64
		fScore[&node] = math.MaxInt64
	}
	gScore[nodeStart] = 0
	fScore[nodeStart] = heuristic(start, end)

	for len(openSet) > 0 {
		current := heap.Pop(&openSet).(*structures.NodesRelations)
		if current.Parent == nodeEnd.Parent {
			return fScore[current]
		}

		var nodeCurr *structures.NodePositions
		filter := bson.M{"parent": current.Parent}
		if err := collectionNodes.FindOne(context.Background(), filter).Decode(&nodeCurr); err != nil {
			return 0
		}

		for _, node := range current.Children {
			var nodeRelation *structures.NodesRelations
			filter := bson.M{"parent": node}
			if err := collection.FindOne(context.Background(), filter).Decode(&nodeRelation); err != nil {
				return 0
			}

			var nodeNeighbor *structures.NodePositions
			filter = bson.M{"parent": node}
			if err := collectionNodes.FindOne(context.Background(), filter).Decode(&nodeNeighbor); err != nil {
				return 0
			}

			tentativeGScore := gScore[current] + heuristic(nodeCurr.Position, nodeNeighbor.Position)
			if tentativeGScore < gScore[current] {
				gScore[nodeRelation] = tentativeGScore
				fScore[nodeRelation] = tentativeGScore + heuristic(nodeNeighbor.Position, end)
				if openSet.Contains(nodeRelation) {
					heap.Push(&openSet, &Item{Node: nodeRelation, Priority: float32(gScore[nodeRelation])})
				}
			}
		}
	}
	return 0
}
